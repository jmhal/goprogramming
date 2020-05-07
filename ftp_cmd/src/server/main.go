package main

import (
   "fmt"
   "os"
   "net"
   "bufio"
   "strings"
   "io"
   "io/ioutil"
   "log"
)

func main() {
   fmt.Println("Servidor FTP.")

   // Abertura passiva.
   conn, err := net.Listen("tcp", ":33000")
   if err != nil {
      log.Fatal(err)
      return
   }
   defer conn.Close()

   // Laço aceitando novas conexões.
   for {
      conn, err := conn.Accept()
      if err != nil {
         log.Fatal(err)
         continue
      }
      go handleRequest(conn)
   }
}

func handleRequest(conn net.Conn) {
   defer conn.Close()
   connReader := bufio.NewReader(conn)
   connWriter := bufio.NewWriter(conn)

   command, _ := connReader.ReadString('\n');
   command = strings.Replace(command, "\n", "", -1)
   arguments := strings.Split(command, " ")

   switch arguments[0] {
   case "list":
      // Listar o diretório
      files, err :=  ioutil.ReadDir(os.Getenv("PWD"));
      if err != nil {
         log.Fatal(err)
         connWriter.WriteString(err.Error() + "\n")
         connWriter.Flush()
         break
      }
      output := ""
      for _, f := range files {
         if f.IsDir() {
            output += f.Name() + "/" + "\n"
         } else {
            output += f.Name() + "\n"
         }
      }

      connWriter.WriteString(output + "$")
      connWriter.Flush()
      break
   case "get":
      // Enviar arquivo requisitado.
      arquivo := arguments[1]
      fmt.Println("Enviando arquivo " + arquivo)
      remoto, err := os.Open(arquivo)
      if err != nil {
         connWriter.WriteString(err.Error() + "\n")
         connWriter.Flush()
         log.Fatal(err)
         break
      }
      _, err = io.Copy(connWriter, remoto)
      if err != nil {
         log.Fatal(err)
         break;
      }
      break
   case "put":
      // Criar o arquivo 
      arquivoRemoto := arguments[1]
      remoto, err := os.Create(arquivoRemoto)
      defer remoto.Close()
      if err != nil {
         log.Fatal(err)
         connWriter.WriteString(err.Error() + "\n")
         connWriter.Flush()
         break
      }

      _, err = io.Copy(remoto, connReader)
      if err != nil {
         log.Fatal(err)
         break
      }
   case "rm":
      arquivoRemoto := arguments[1]
      err := os.Remove(arquivoRemoto)
      if err != nil {
         log.Fatal(err)
         connWriter.WriteString(err.Error() + "$")
         connWriter.Flush()
         break
      }
      connWriter.WriteString("Arquivo " + arquivoRemoto + " removido.$")
      connWriter.Flush()
      break
   }
}
