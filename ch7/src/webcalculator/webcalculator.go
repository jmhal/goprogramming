package main

import (
  "fmt"
  "net/http"
  "html/template"
  "log"
  "eval"
)

const index_html = `<html>
   <head>
      <style>
      table, th, td {
         border: 1px solid black;
      }
      </style>
      <title>
         Web Calculator 
      </title>
   </head>
   <body>
      {{if .Success}}
      <h1> {{.Result}} </h1> 
      {{else}}
      <h1> Web Calculator </h1>
      <form method="POST">
         <label> Expression: </label><br/>
	 <input type="text" name="expression"><br/>
	 <input type="submit">
      </form>
      {{end}}
   </body>
</html>
`

func compute(exprParameter string) string {
   expr, err := eval.Parse(exprParameter)
   if err != nil {
      log.Println(err)
   }
   env := eval.Env{}
   expr.Environment(env)
   return  fmt.Sprintf("%.6g", expr.Eval(env))
}

func main() {
   tmpl := template.Must(template.New("index").Parse(index_html))
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      if r.Method != http.MethodPost {
         tmpl.Execute(w, nil)
	 return
      }
      expression := r.FormValue("expression")
      value := compute(expression)
      data := struct {
         Success bool
	 Result string
      } {
         true,
	 value,
      }
      tmpl.Execute(w, data)
   })
   log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
