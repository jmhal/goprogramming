package main

import (
   "os"
   "fmt"
   "strings"
   "net/http"
   "golang.org/x/net/html"
)

func main() {
   url := os.Args[1]
   words, images, err := CountWordsAndImages(url)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(url + " has:")
   fmt.Printf("%d words, and %d images.\n", words, images)

   return
}

func CountWordsAndImages(url string) (words, images int, err error) {
   resp, err := http.Get(url)
   if err != nil {
      fmt.Println("Unable to fetch: " + url)
      return
   }
   doc, err := html.Parse(resp.Body)
   resp.Body.Close()
   if err != nil {
      err = fmt.Errorf("parsing HTML: %s", err)
      return
   }
   words, images = countWordsAndImages(doc)
   return
}

func countWordsAndImages(n *html.Node) (words, images int) {
   if n != nil {
      if n.Type == html.ElementNode {
	 if n.Data == "img" {
	   images++
         }
	 words += len(strings.Fields(n.Data))
         for _, a := range n.Attr {
	    words += len(strings.Fields(a.Val))
	 }
      }

      if n.NextSibling != nil {
         siblingWords, siblingImages := countWordsAndImages(n.NextSibling)
	 words += siblingWords
	 images += siblingImages
      }
      if n.FirstChild != nil {
         firstChildWords, firstChildImages := countWordsAndImages(n.FirstChild)
	 words += firstChildWords
	 images += firstChildImages
      }
   }
   return
}

