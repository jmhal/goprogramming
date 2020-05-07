// Fetch all busca URLs em paralelo e informe os tempos gastos e os tamanhps. 
package main

import (
   "fmt"
   "io"
   "io/ioutil"
   "net/http"
   "os"
   "time"
)

func main() {
   start := time.Now()
   ch := make(chan string)
   
   fileName := os.Args[1]
   f, err := os.Create(fileName)
   if err != nil {
      fmt.Fprintf(os.Stderr, "Cannot open file: %s", fileName)
      os.Exit(1)
   }

   for _, url := range os.Args[2:] {
      go fetch(url, ch) // inicia uma gorrotina
   }
   for range os.Args[2:] {
      fmt.Fprintf(f, <-ch) // recebe do canal ch
   }
   fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string) {
   start := time.Now()
   resp, err := http.Get(url)

   if err != nil {
     ch <- fmt.Sprint(err) // envia para o canal ch
     return
   }
   nbytes, err := io.Copy(ioutil.Discard, resp.Body)
   resp.Body.Close() // evita vazamento de recursos

   if err != nil {
      ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
      return
   }
   
   secs := time.Since(start).Seconds()
   ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

