// forms.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		nome := r.FormValue("nome")
		sobrenome := r.FormValue("sobrenome")

		file, err := os.OpenFile("confirmacoes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Impossível abrir arquivo.", err)
		}
		defer file.Close()

		fmt.Fprintf(file, "%s %s\n", nome, sobrenome)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
