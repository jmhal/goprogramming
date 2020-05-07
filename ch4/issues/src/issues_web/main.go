package main

import (
   "log"
   "time"
   "net/http"
   "html/template"
   "github"
)

const index_html = `<html>
<head><title>{{.PageTitle}}</title></head>
<body>
<h1> {{.PageHeader}}
<hr>

<h2>
<form action="/submit" method="post">
   <label> Busca: </label>
   <input type="text" name="busca"> <br>
   <input type="radio" name="searchtype" value="bugs" checked> Bugs<br>
   <input type="radio" name="searchtype" value="milestones"> Milestones<br>
   <input type="radio" name="searchtype" value="users"> Users<br>
   <input type="submit" value="Buscar"> <br>
</form>
<hr>

</body>
</html>
`

const submit_html = `<html>
<head>
<title>{{.PageTitle}}</title>
<style>
table, th, td {
    border: 1px solid black;
}
</style>
</head>
<body>
<h1> Informações de {{.PageHeader}}.
<hr>
<h2>

{{.Results.TotalCount}} resultados:
<table style="width:100%">
<tr>
   <th>Número</th>
   <th>Título</th>
   <th>Estado</th>
   <th>Usuário</th>
   <th>Data de Criação</th>
   <th>Conteúdo</th>
</tr>
{{range .Results.Items}}
<tr>
   <th>{{.Number}}</th>
   <th>{{.Title | printf "%.64s"}}</th>
   <th>{{.State}}</th>
   <th>{{.User.Login}}</th>
   <th>{{.CreatedAt | daysAgo}}</th>
   <th>{{.Body}}</th>
</tr>
{{end}}
</table> 
</body>
</html>
`

func main(){
   // Redirecionar a raiz para página que carrega o template do index
   http.HandleFunc("/", index)

   // Redirecionar o caminho /submit para receber a imagem e armazenar no diretório
   http.HandleFunc("/submit", submit)

   // Ativa o servidor
   log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func daysAgo(t time.Time) int {
  return int(time.Since(t).Hours() / 24)
}

func index(w http.ResponseWriter, r *http.Request) {
   data := struct {
      PageTitle string
      PageHeader string
   }{
      "Informações do GitHub",
      "Forneça os termos de busca:",
   }
   tmpl := template.Must(template.New("index").Parse(index_html))
   tmpl.Execute(w, data)
}

func submit(w http.ResponseWriter, r *http.Request) {
   r.ParseForm()
   // log.Println(r)

   searchdata := r.Form["busca"][0]
   searchtype := r.Form["searchtype"][0]

   var searchstring []string
   if searchtype == "bugs" {
      searchstring = []string{"label:bug ", searchdata}
   } else if searchtype == "milestones" {
      searchstring = []string{"milestone:", searchdata}
   } else if searchtype == "users" {
      searchstring = []string{"author:" + searchdata, ""}
   }
   log.Println(searchstring)
   results, err := github.SearchIssues(searchstring)
   if err != nil {
      log.Fatal(err)
   }
   // log.Println(results)

   data := struct {
      PageTitle string
      PageHeader string
      Results *github.IssuesSearchResult
   }{
      "Informações do GitHub",
      "Resultado da busca: " + searchstring[0] + searchstring[1],
      results,
   }
   tmpl := template.Must(template.New("submit").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(submit_html))
   tmpl.Execute(w, data)
}
