// Fetch exibe o conteúdo encotrado em cada URL especificada.
package main

import (
   "fmt"
   "io"
   "strings"
   "net/http"
   "os"
)

func main() {
   for _, url := range os.Args[1:] {
      if !strings.HasPrefix(url, "http://") {
         url = "http://" + url
      }

      resp, err := http.Get(url)
      if err != nil {
         fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	 os.Exit(1)
      }
      _, err = io.Copy(os.Stdout, resp.Body)
      
      if err != nil {
         fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	 os.Exit(1)
      }
      fmt.Fprintf(os.Stderr, "http status: %s\n", resp.Status)
   }
}
