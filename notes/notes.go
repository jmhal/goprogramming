// notes.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jmhal/goprogramming/notes/notesdb"
)

var data []notesdb.Note

func list_notes(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("listar_notas.html"))

	// Aqui teria que recuperar as notas do MongoDB
	data = notesdb.GetNotes()

	tmpl.Execute(w, struct{ Notes []notesdb.Note }{data})
	return
}

func add_notes(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("adicionar_nota.html"))
	tmpl.Execute(w, struct{ Notes []notesdb.Note }{data})
	return
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		titulo := r.FormValue("titulo")
		conteudo := r.FormValue("conteudo")

		// Aqui teria que adicionar uma nota ao MongoDB
		notesdb.InsertNote(titulo, conteudo)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
	return
}

func main() {
	mongouri := strings.Replace(os.Getenv("MONGO_PORT"), "tcp", "mongodb", 1)
	fmt.Println(mongouri)
	//mongouri := "mongodb://172.17.0.2:27017"
	notesdb.CreateConnection(mongouri)

	http.HandleFunc("/add", add)
	http.HandleFunc("/list-notes", list_notes)
	http.HandleFunc("/add-notes", add_notes)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
