package main

import (
   "fmt"
   "os"
   "bufio"
   "strings"
   "net"
)


func main() {
   fmt.Println("FTP Client.")

   if (len(os.Args) < 3) {
      fmt.Println("usage: ./client <IP> <PORT>")
      return
   }

   // Estabelece a conexão
   IP := os.Args[1]
   PORT := os.Args[2]
   conn, err := net.Dial("tcp", IP + ":" + PORT)
   defer conn.Close()
   connReader := bufio.NewReader(conn)
   if err != nil {
      fmt.Println(err)
      return
   }

   // Variável para armazenar o comando
   reader := bufio.NewReader(os.Stdin)
   var command string
   for  {
      fmt.Print("-> ")

      // Ler o comando e retirar o nova linha do final.
      command, _ = reader.ReadString('\n')
      command = strings.Replace(command, "\n", "", -1)

      // Separar os argumentos
      arguments := strings.Split(command, " ")

      // Tratar os possíveis comandos.
      switch arguments[0] {
      case "quit":
         fmt.Println("bye...")
         fmt.Fprintf(conn, "quit\n")
         conn.Close()
         return;
      case "list":
         var dir string
         if len(arguments) < 2 {
            dir = ""
         } else {
            dir = arguments[1]
         }
         fmt.Fprintf(conn, "list:" + dir + "\n")

         output, err := connReader.ReadString('\n')
         if err != nil {
            fmt.Println(err)
            fmt.Fprintf(conn, "quit\n")
            conn.Close()
            return
         }

         files := strings.Split(output, ":")
         files = files[:len(files)-1]
         for _, file := range files {
            fmt.Println(file)
         }

         break;
      case "mkdir":
         break;
      case "cd":
         var dir string
         if  len(arguments) < 2 {
            fmt.Println("usage: cd <dir>")
            break;
         } else {
            dir = arguments[1]
         }
         fmt.Fprintf(conn, "cd:" + dir + "\n")

         output, err := connReader.ReadString('\n')
         if err != nil {
            fmt.Println(err)
            fmt.Fprintf(conn, "quit\n")
            conn.Close()
            return
         }

         fmt.Println(output)
         break;
      case "get":
         break;
      case "put":
         break;
      case "rm":
         break;
      case "rmdir":
         break;
      }
   }
}
