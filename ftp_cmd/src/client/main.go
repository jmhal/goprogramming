package main

import (
   "fmt"
   "os"
   "log"
   "net"
   "bufio"
   "strings"
   "io"
)


func main() {
   fmt.Println("Cliente FTP.")
   if len(os.Args) < 4 {
      fmt.Println("Uso:")
      fmt.Println("   ./cliente <IP> <PORTA> list")
      fmt.Println("      listar os arquivos do servidor." )
      fmt.Println("   ./cliente <IP> <PORTA> get <ARQUIVOREMOTO> <ARQUIVOLOCAL>")
      fmt.Println("      copia <ARQUIVOREMOTO> do servidor para o arquivo <ARQUIVOLOCAL>." )
      fmt.Println("   ./cliente <IP> <PORTA> put <ARQUIVOLOCAL> <ARQUIVOREMOTO>")
      fmt.Println("      copia <ARQUIVOLOCAL> para <ARQUIVOREMOTO> no servidor." )
      fmt.Println("   ./cliente <IP> <PORTA> rm <ARQUIVOREMOTO>")
      fmt.Println("      remove <ARQUIVOREMOTO> do servidor." )
      return
   }

   // Cria a Conexão
   serverIP := os.Args[1]
   serverPORT := os.Args[2]
   conn, err := net.Dial("tcp", serverIP + ":" + serverPORT)
   defer conn.Close()
   if err != nil {
      log.Fatal(err)
      return
   }
   connReader := bufio.NewReader(conn)
   connWriter := bufio.NewWriter(conn)

   arguments := os.Args[3:]
   switch arguments[0] {
   case "list":
      connWriter.WriteString("list\n") 
      err := connWriter.Flush()
      if err != nil {
         log.Fatal(err)
         break
      }

      output, err := connReader.ReadString('$')
      if err != nil {
         log.Fatal(err)
         break
      }
      fmt.Println("Conteúdo do Diretório Remoto:")
      fmt.Println(strings.Trim(output, "$"))
      break
   case "get":
      arquivoRemoto := arguments[1]
      arquivoLocal := arguments[2]

      // Envia qual arquivo é desejado.
      connWriter.WriteString("get " + arquivoRemoto + "\n")
      err := connWriter.Flush()
      if err != nil {
         log.Fatal(err)
         break
      }

      // Cria o arquivo local.
      local, err := os.Create(arquivoLocal)
      if err != nil {
         log.Fatal(err)
         break
      }
      defer local.Close()

      _, err = io.Copy(local, connReader)
      if err != nil {
         log.Fatal(err)
         break
      }
      break
   case  "put":
      arquivoLocal := arguments[1]
      arquivoRemoto := arguments[2]

      // Envia qual arquivo a ser criado.
      connWriter.WriteString("put " + arquivoRemoto + "\n")
      err := connWriter.Flush()
      if err != nil {
         log.Fatal(err)
         break
      }

      // Abre o arquivo local.
      local, err := os.Open(arquivoLocal)
      if err != nil {
         log.Fatal(err)
         break
      }
      defer local.Close()

      // Faz a cópia.
      _, err = io.Copy(connWriter, local)
      if err != nil {
         log.Fatal(err)
         break
      }
      break
   case "rm":
      arquivoRemoto := arguments[1]

      // Envia arquivo a ser removido. 
      connWriter.WriteString("rm " + arquivoRemoto + "\n")
      err := connWriter.Flush()
      if err != nil {
         log.Fatal(err)
         break
      }

      // Recebe a mensagem de sucesso ou fracasso.
      output, err := connReader.ReadString('$')
      if err != nil {
         log.Fatal(err)
         break
      }
      fmt.Println(strings.Trim(output, "$"))
      break
   }
}
