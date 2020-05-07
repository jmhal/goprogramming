package main

import (
   "fmt"
   "io"
   "strings"
)

type LimitedReader struct {
   limitBytes *int64
   readBytes *int64
   internalReader io.Reader
}

func (l LimitedReader) Read(p []byte) (int, error) {
   bytesToRead := int64(len(p))
   bytesAvailable := *(l.limitBytes) - *(l.readBytes)
   if (bytesToRead <= bytesAvailable) {
      n, err := l.internalReader.Read(p)
      *(l.readBytes) += int64(n)
      return n, err
   }
   data := make([]byte, bytesToRead)
   n, _ := l.internalReader.Read(data)
   copy(p[:], data)
   return n, io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
   var l LimitedReader
   var counter, limit int64
   counter = 0
   limit = n
   l.limitBytes = &limit
   l.readBytes = &counter
   l.internalReader = r
   return l
}

func main() {
   sr := strings.NewReader("João Marcelo")
   nr := LimitReader(sr, 5)
   p := make([]byte, 2) 
   n, err := nr.Read(p)
   if err != nil {
      fmt.Println("Error reading first two bytes.")
      return
   }
   fmt.Printf("%d %s \n", n, p)
   n, err = nr.Read(p)
   if err != nil {
      fmt.Println("Error reading following two bytes.")
      return
   }
   fmt.Printf("%d %s \n", n, p)
   n, err = nr.Read(p)
   if err != nil {
      fmt.Println("Error reading following following two bytes.")
      fmt.Printf("%s \n", p)
      return
   }
   fmt.Printf("%d %s \n", n, p)
   return
}
