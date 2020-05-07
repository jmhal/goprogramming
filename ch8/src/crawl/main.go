package main

import (
   "os"
   "fmt"
   "links"
   "strings"
   "net/http"
   "log"
   "bytes"
   "strconv"
)

// DataType with a url and depth
type WorkItem struct {
   url string
   level int
}

// Global variables
var domains []string
var currentDir string
var done = make(chan struct{})

// Verifies if a string belongs to a map or not
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

   list, err := links.Extract(url, done)
   if err != nil {
      log.Print(err)
   }

   list = append(list, url)
   log.Printf("Processing link list: %s", list)
   filteredList := []string{}

   for _, linkUrl := range list {
      log.Printf("Current Link: %s", linkUrl)
      belongsToDomain := false
      for _, domain := range domains {
         if strings.Contains(linkUrl, domain) {
            log.Printf("Link %s has domain %s", linkUrl , domain)
            belongsToDomain = true
         }
      }
      if belongsToDomain {
         // Deals with # links
         link := strings.TrimSpace(strings.Join(strings.Split(linkUrl, "#")[:1],""))
         if linkAlreadyCovered(link, filteredList) {
	         continue
         }

         // There are two kinds of links
         // 1- http://example.com/test/ or http://example.com/test create the directory test and place the page inside index.html
         // 2- http://example.com/path/test.html create the directory path and place test.html

         extensions := []string{".html", ".htm", ".css", ".php", ".png", ".jpeg", ".jpg", ".png", ".gif"}
         hasExtension := false
         for _, extension := range extensions {
            if strings.HasSuffix(link, extension) {
               log.Printf("Link %s has extension %s", link, extension)
               hasExtension = true
            }
         }

         path := strings.TrimPrefix(strings.TrimPrefix(link, "https://"), "http://")
         pathElements := strings.Split(path,"/")

         baseUrl := strings.Join(strings.Split(link, "/")[0:3], "/")
         log.Printf("The base url for link %s is %s", link, baseUrl)
         baseDomain := strings.Split(link,"/")[2]
         log.Printf("The base domain for link %s is %s", link, baseDomain)

         if hasExtension {
            dir := strings.Join(pathElements[0:len(pathElements) - 1], "/")

            // Create Directory
            log.Printf("Creating dir: %s", dir)
            err1 := os.MkdirAll(dir, 0755)
            if err1 != nil {
               log.Printf("Failed to create dir: %s\n", dir)
            }

            // Fetch the page
	         resp, err2 := http.Get(link)
            if err2 != nil {
               log.Println(err2)
            }
            if resp.StatusCode != http.StatusOK {
               resp.Body.Close()
               log.Printf("getting %s: %s", url, resp.Status)
            }

            // Create the file
            log.Printf("Creating file %s for link %s", path, link)
            f, err := os.Create(path)
            defer f.Close()
            if err != nil {
	            log.Printf("cannot write file: %s", path)
            }
            buf := new(bytes.Buffer)
            buf.ReadFrom(resp.Body)

            page := buf.String()
            page = strings.Replace(page, baseUrl, currentDir + "/" + baseDomain, -1)

	         f.Write([]byte(page))

            // Closes the request
            resp.Body.Close()
         } else {
            // Create Directory
            log.Printf("Creating dir: %s", path)
            err1 := os.MkdirAll(path, 0755)
            if err1 != nil {
               log.Printf("Failed to create dir: %s\n", path)
            }

            // Fetch the page
	         resp, err2 := http.Get(link)
            if err2 != nil {
               log.Println(err2)
            }
            if resp.StatusCode != http.StatusOK {
               resp.Body.Close()
              log.Printf("getting %s: %s", url, resp.Status)
            }

            // Create the file
            log.Printf("Creating file %s for link %s", path + "/index.html", link)
            f, err := os.Create(path + "/index.html")
            defer f.Close()
            if err != nil {
	            log.Printf("cannot write file: %s", path + "/index.html")
            }
            buf := new(bytes.Buffer)
            buf.ReadFrom(resp.Body)

            page := buf.String()
            page = strings.Replace(page, baseUrl, currentDir + "/" + baseDomain, -1)

	         f.Write([]byte(page))

            // Closes the request
            resp.Body.Close()
         }

         // Mark the link as visited. 
         filteredList = append(filteredList, link)
      } else {
         log.Printf("LINK %s does not belong to domains.", linkUrl)
      }
   }
   return filteredList
}

func main() {
   // Get current working directory
   pwd, errDir := os.Getwd()
   if errDir != nil {
      log.Fatal(errDir)
   }
   currentDir = pwd
   log.Printf("Current working dir: %s", currentDir)

   urls := []WorkItem{}
   var depth int
   for _, param := range os.Args[1:] {
      if strings.Contains(param, "-depth=") {
         d, err := strconv.Atoi(strings.Split(param, "=")[1])
         if err != nil {
            log.Println("Wrong depth!!!")
            os.Exit(1)
         }
         depth = d
      } else {
         domain := strings.TrimPrefix(param, "https://")
         domain = strings.TrimPrefix(domain, "http://")
         domain = strings.TrimSuffix(domain, "/")
         domains = append(domains, domain)
         urls = append(urls, WorkItem{param, 0} )
      }
   }

   log.Printf("Domains: %s\n", domains)
   log.Printf("Defined depth: %d", depth)

   worklist := make(chan []WorkItem)
   unseenLinks := make(chan WorkItem)

   go func() { worklist <- urls }()

   for i := 0; i < 5; i++ {
      go func() {
         for link := range unseenLinks {
            foundLinks := crawl(link.url)
            foundItems := []WorkItem{}
            for _, url := range foundLinks {
               foundItems = append(foundItems, WorkItem{url, link.level + 1})
            }
            go func() { worklist <- foundItems }()
         }
      }()
   }

   go func() {
      os.Stdin.Read(make([]byte,1))
      close(done)
   }()

   seen := make(map[string]bool)
   for {
      select {
      case <- done:
         return
      case itemsList := <-worklist:
         for _, item := range itemsList {
            if item.level >= depth {
               continue
            }
            if !seen[item.url] {
               seen[item.url] = true
               unseenLinks <- item
            }
         }
      }
   }
}
