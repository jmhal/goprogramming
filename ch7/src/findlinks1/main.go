package main

import (
	"fmt"
	"os"
        "io"
	"golang.org/x/net/html"
)

type BufferString struct {
   str string
   read *bool
}

func (b BufferString) Read(p []byte) (int, error) {
   if *(b.read) {
      copy(p[:], string(b.str))
      *(b.read) = false
      return len(string(b.str)), nil
   } else {
      return 0, io.EOF 
   }
}

func NewReader(str string) (io.Reader) {
   var b BufferString
   var status bool
   status = true
   b.str = str
   b.read = &status
   return b
}

func main() {
   page := "<html><head><title>Teste</title></head><body>Isto é um teste.</body></html>"

   doc, err := html.Parse(NewReader(page))
   if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
      os.Exit(1)
   }
   fmt.Println(doc)
}


