package main

import (
   "encoding/xml"
   "fmt"
   "io"
   "os"
   "strings"
)

func main() {
   dec := xml.NewDecoder(os.Stdin)
   var stack []xml.StartElement // []string
   for {
      tok, err := dec.Token()
      if err == io.EOF {
         break
      } else if err != nil {
         fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
	 os.Exit(1)
      }
      switch tok := tok.(type) {
         case xml.StartElement :
	    stack = append(stack, tok)
	 case xml.EndElement :
	    stack = stack[:len(stack)-1]
	 case xml.CharData:
	    if containsMatch(stack, os.Args[1:]) {
	       elements := []string{}
	       for _, e := range stack {
	          elements = append(elements, e.Name.Local)
	       }
	       fmt.Printf("%s: %s\n", strings.Join(elements, " "), tok)
	    }
      }
   }
}

func containsAll(x, y []string) bool {
   for len(y) <= len(x) {
      if len(y) == 0 {
         return true
      }
      if x[0] == y[0] {
         y = y[1:]
      }
      x = x[1:]
   }
   return false
}

func containsMatch(x []xml.StartElement, y []string) bool {
   for len(y) <= len (x) {
      if len(y) == 0 {
         return true
      }
      if x[0].Name.Local == y[0] || contains(x[0].Attr, y[0]) {
         y = y[1:]
      }
      x = x[1:]
   }
   return false
}

func contains(x []xml.Attr, y string) bool {
   for _, v := range x {
      if v.Name.Local == y {
         return true
      }
   }
   return false
}
