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
   forEachNode(doc, outline("pre"), outline("post"))
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
   /*
   if n.Type == html.ElementNode && n.FirstChild == nil {
      var attributes string
      for _, attr := range n.Attr {
         attributes += " " + attr.Key + "=" +  attr.Val
      }
      fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data + attributes)
      return
   }
   */
   if pre != nil {
      pre(n)
   }

   for c := n.FirstChild; c != nil; c = c.NextSibling {
      forEachNode(c, pre, post)
   }

   if post != nil {
      post(n)
   }
}

func outline(functionType string) func(n *html.Node) {
   var depth int
   if functionType == "pre" {
      return func(n *html.Node) {
         if n.Type == html.ElementNode {
            var attributes string
            for _, attr := range n.Attr {
               attributes += " " + attr.Key + "=" +  attr.Val
            }
            fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data + attributes)
            depth++
         } else if n.Type == html.TextNode {
            data := strings.TrimSpace(n.Data)
            if data != "" {
               fmt.Printf("%*s %s\n", depth * 2, "", data)
            }
         }
      }
   } else {
      return func(n *html.Node) {
         if n.Type == html.ElementNode {
            depth--
            fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
         }
      }
   }
}
/*
func startElement(n *html.Node) {
   if n.Type == html.ElementNode {
      var attributes string
      for _, attr := range n.Attr {
         attributes += " " + attr.Key + "=" +  attr.Val
      }
      fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data + attributes)
      depth++
   } else if n.Type == html.TextNode {
      data := strings.TrimSpace(n.Data)
      if data != "" {
         fmt.Printf("%*s %s\n", depth * 2, "", data)
      }
   }
}

func endElement(n *html.Node) {
   if n.Type == html.ElementNode {
      depth--
      fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
   }
}
*/
