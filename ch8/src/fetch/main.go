package main

import (
   "os"
   "fmt"
   "io"
   "net/http"
   "path"
)

type Response struct {
   fileName string
   nBytes int64
   err error
}

func fetch(urls []string) Response {
   responses := make(chan Response, len(urls))
   for _, url := range urls {
      go func (url string) {
         fmt.Printf("recovering %s\n", url)
         resp, err := http.Get(url)
         if err != nil {
            responses <- Response{
               fileName: "",
               nBytes: 0,
               err: err,
            }
         }
         defer resp.Body.Close()

         local := path.Base(resp.Request.URL.Path)
         if (local == "/") || (local == ".") {
            local = "index.html"
         }

         f, err := os.Create(local)
         defer f.Close()
         if err != nil {
            responses <- Response{
               fileName: "",
               nBytes: 0,
               err: err,
            }
         }

         n, err := io.Copy(f, resp.Body)
         responses <- Response{
            fileName: local,
            nBytes: n,
            err: err,
         }
      }(url)
   }
   return <- responses
}

func main() {
   response := fetch(os.Args[1:])
   if response.err == nil {
      fmt.Println(response.fileName)
   } else {
      fmt.Println(response.err)
   }
}
