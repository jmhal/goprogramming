package main

import (
   "fmt"
   "os"
   "golang.org/x/net/html"
)

type HTML struct {
   Links []string
   Scripts []string
   Imgs []string
}

func main() {
   doc, err := html.Parse(os.Stdin)
   if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
      os.Exit(1)
   }
   var elements HTML
   elements = visit(elements, doc)

   fmt.Println("LINKS:")
   for _, link := range elements.Links {
      fmt.Println(link)
   }

   fmt.Println("SCRIPTS:")
   for _, script := range elements.Scripts {
      fmt.Println(script)
   }

   fmt.Println("IMAGES:")
   for _, img := range elements.Imgs {
      fmt.Println(img)
   }

   /*
   m := make(map[string]int)
   elementsNumbers(m, doc)
   for element, number := range m {
      fmt.Printf("%s: %d\n", element, number)
   }

   printTextElements(doc)
   */
}

func elementsNumbers(elements map[string]int, n *html.Node) {
   if n != nil {
      if n.Type == html.ElementNode {
         elements[n.Data]++;
      }
      if n.NextSibling != nil {
         elementsNumbers(elements, n.NextSibling)
      }
      if n.FirstChild != nil {
         elementsNumbers(elements, n.FirstChild)
      }
   }
   return
}

func printTextElements(n *html.Node) {
   if n != nil {
      if n.Type == html.ElementNode {
         if n.Data != "script" && n.Data != "style" && n.Data != "stylesheet" {
	    for _, a := range n.Attr {
	       fmt.Println(a.Val)
	    }
	 }
      }
      if n.NextSibling != nil {
         printTextElements(n.NextSibling)
      }
      if n.FirstChild != nil {
         printTextElements(n.FirstChild)
      }
   }
   return
}

func visit(elements HTML, n *html.Node) HTML {
   if n != nil {
      if n.Type == html.ElementNode {
         if n.Data == "a" {
            for _, a := range n.Attr {
               if a.Key == "href" {
	          elements.Links = append(elements.Links, a.Val)
	       }
            }
         } else if n.Data == "script" {
	    for _, a := range n.Attr {
	       if a.Key == "src" {
	          elements.Scripts = append(elements.Scripts, a.Val)
	       }
	    }
	 } else if n.Data == "img" {
	    for _, a := range n.Attr {
	       if a.Key == "src" {
	          elements.Imgs = append(elements.Imgs, a.Val)
	       }
	    }
         }
      }
      if n.NextSibling != nil {
         elements = visit(elements, n.NextSibling)
      }
      if n.FirstChild != nil {
         elements = visit(elements, n.FirstChild)
      }
   }
   return elements
}
