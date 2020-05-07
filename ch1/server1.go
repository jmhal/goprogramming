// Server1 é um servidor de "eco" mínimo. 
package main

import (
   "fmt"
   "log"
   "net/http"
)

func main() {
   http.HandleFunc("/", handler)  // cada requisição chama handler
   log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
