package main

import (
   "log"
   "net"
   "os"
   "io"
)

func main() {
   conn, err := net.Dial("tcp", "localhost:8000")
   if err != nil {
      log.Fatal(err)
   }
   done := make(chan struct{})

   go func() {
      _, err := io.Copy(os.Stdout, conn)
      if err != nil {
         log.Println(err)
      }
      log.Println("done")
      done <-struct{}{}
   }()

   mustCopy(conn, os.Stdin)

   conn.(*net.TCPConn).CloseWrite()
   <-done
   log.Print("finish.")
}

func mustCopy(dst io.Writer, src io.Reader) {
   if _, err := io.Copy(dst, src); err != nil {
      log.Println(err)
   }
}

