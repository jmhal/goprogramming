package main

import (
   "fmt"
   "links"
   "log"
   "os"
   "strings"
   "net/http"
   "bytes"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item. 
func breadthFirst(f func(item string) []string, worklist []string) {
   seen := make(map[string]bool)
   for len(worklist) > 0 {
      items := worklist
      worklist = nil
      for _, item := range items {
         if !seen[item] {
	    seen[item] = true
	    worklist = append(worklist, f(item)...)
	 }
      }
   }
}

func linkAlreadyCovered(link string, list []string) bool {
   for _, a := range list {
      if link == a {
         return true
      }
   }
   return false
}

func crawl(url string) []string {
   fmt.Println(url)

   list, err := links.Extract(url)
   if err != nil {
      log.Print(err)
   }
   filteredList := []string{}
   for _, a := range list {
      if strings.Contains(a, url) {
         // Deals with # links
         link := strings.Join(strings.Split(a, "#")[:1],"")
         if linkAlreadyCovered(link, filteredList){
	    continue
	 }

         // Create directory
         dir := strings.TrimPrefix(link, "https://")
         dir = strings.TrimPrefix(dir, "http://")
         err := os.MkdirAll(dir, 0755)
         if err != nil {
            log.Printf("Failed to create dir: %s\n", dir)
         }
	 // Download page to the directory
	 resp, err := http.Get(url)
         if err != nil {
            log.Println(err)
         }
         if resp.StatusCode != http.StatusOK {
            resp.Body.Close()
            log.Printf("getting %s: %s", url, resp.Status)
         }

	 f, err := os.Create(dir + "/index.html")
	 defer f.Close()
	 if err != nil {
	    log.Printf("cannot write file: %s", dir + "/index.html")
	 }
         buf := new(bytes.Buffer)
         buf.ReadFrom(resp.Body)
	 f.Write(buf.Bytes())

         resp.Body.Close()
	 // Returns the list of links
         filteredList = append(filteredList, link)
      }
   }
   return filteredList
}

func main() {
   breadthFirst(crawl, os.Args[1:])
}
