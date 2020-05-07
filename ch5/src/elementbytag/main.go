package main

import (
   "fmt"
   "os"
   "golang.org/x/net/html"
   "net/http"
   "strings"
)

func main() {
   resp, err := http.Get(os.Args[1])
   if err != nil {
       fmt.Fprintf(os.Stderr, "recuperar: %v\n", err)
      os.Exit(1)
   }

   doc, err := html.Parse(resp.Body)
   if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
      os.Exit(1)
   }
   img := ElementsByTagName(doc, "img", "svg")
   for _, i := range img {
      fmt.Println(i.Data)
   }
   headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4", "br")
   for _, h := range headings {
      fmt.Println(h.Data)
   }

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
   elements := []*html.Node{}

   if doc.Type == html.ElementNode {
      for _, s := range name {
         if strings.Contains(doc.Data, s) {
	    elements = append(elements, doc)
	    break
	 }
      }
   }
   for c := doc.FirstChild; c != nil; c = c.NextSibling {
      elements = append(elements, ElementsByTagName(c, name...)...)
   }
   return elements
}
