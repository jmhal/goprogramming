package main

import (
   "fmt"
   "os"
   "net"
   "bufio"
   "strings"
   "io/ioutil"
)

func main() {
   fmt.Println("Servidor FTP.")

   // Abertura passiva.
   conn, err := net.Listen("tcp", ":33000")
   if err != nil {
      fmt.Println("Erro na abertura da conexão:", err.Error())
      os.Exit(1)
   }
   defer conn.Close()

   // Laço aceitando novas conexões.
   for {
      conn, err := conn.Accept()
      if err != nil {
         fmt.Println("Erro ao aceitar conexão:", err.Error())
      }
      go handleRequest(conn)
   }
}

func handleRequest(conn net.Conn) {
   defer conn.Close()
   currentDir := os.Getenv("PWD")
   fmt.Println("Current Dir: " + currentDir)
   connReader := bufio.NewReader(conn)
   for {
      command, _ := connReader.ReadString('\n');
      command = strings.Replace(command, "\n", "", -1)
      arguments := strings.Split(command, ":")
      switch arguments[0] {
      case "quit":
         fmt.Println("closing connection...")
         conn.Close()
         return
      case "list":
         var dir string
         if arguments[1] == "" {
            dir = currentDir
         } else {
            dir = arguments[1]
         }
         fmt.Println("list dir: " + dir)

         // Listar o diretório
         files, err :=  ioutil.ReadDir(dir);
         if err != nil {
            fmt.Fprintf(conn, err.Error() + "\n")
            break
         }
         output := ""
         for _, f := range files {
            if f.IsDir() {
               output += f.Name() + "/" + ":"
            } else {
               output += f.Name() + ":"
            }
            fmt.Println(f.Name())
         }

         fmt.Fprintf(conn, output + "\n")
         break
      case "cd" :
         dir := arguments[1]
         fmt.Println("Target dir: " + dir)
         err := os.Chdir(dir)
         if err != nil {
            fmt.Println(err)
            fmt.Fprintf(conn, err.Error() + "\n")
            break
         }
         currentDir, err = os.Getwd()
         fmt.Println("Current Dir: " + currentDir)
         fmt.Fprintf(conn, currentDir + "\n")
         break;
      }
   }
}
