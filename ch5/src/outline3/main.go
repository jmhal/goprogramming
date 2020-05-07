package main

import (
   "fmt"
   "os"
   "golang.org/x/net/html"
   "net/http"
)

var depth int

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
   ElementById(doc, os.Args[2])
   
}

func ElementById (doc *html.Node, id string) *html.Node {
   return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string)(bool), post func(n *html.Node)) (*html.Node) {
   status := true
   if pre != nil {
      status = pre(n, id)
   }
   if !status {
      return n
   }

   for c := n.FirstChild; c != nil; c = c.NextSibling {
      nodeFound := forEachNode(c, id, pre, post)
      if nodeFound != nil {
         return nodeFound
      }
   }

   if post != nil {
      post(n)
   }

   return nil
}

func startElement(n *html.Node, id string) (bool) {
   if n.Type == html.ElementNode {
      fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
      depth++
      for _, a := range n.Attr {
         if a.Key == id {
	    return false
	 }
      }
   }
   return true
}

func endElement(n *html.Node) {
   if n.Type == html.ElementNode {
      depth--
      fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
   }
}
