package main

import (
  "fmt"
  "net/http"
  "html/template"
  "strconv"
  "log"
)

type dollars float32
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

const list_html = `<html>
   <head>
      <style>
      table, th, td {
         border: 1px solid black;
      }
      </style>
      <title>
         Items and Prices
      </title>
   </head>
   <body>
      <table>
         <tr>
	    <th> Item  </th>
	    <th> Price </th>
	 </tr>
	 {{ range $item, $price := .DB}}
	 <tr> 
	    <td> {{$item}} </td>
	    <td> {{$price}} </td>
	 </tr>
	 {{end}}
      </table>
   </body>
</html>
`

func (db database) list (w http.ResponseWriter, req *http.Request) {
   /*
   for item, price := range db {
      fmt.Fprintf(w, "%s: %s\n", item, price)
   }
   */
   data := struct {
      DB database
   }{
      db,
   }
   tmpl := template.Must(template.New("index").Parse(list_html))
   tmpl.Execute(w, data)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
   item := req.URL.Query().Get("item")
   price := req.URL.Query().Get("price")
   pricevalue, err := strconv.ParseFloat(price,32)
   if err != nil {
      fmt.Fprintf(w, "invalid price: %q\n", price)
      return
   }
   db[item] = dollars(pricevalue)
   return
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
   item := req.URL.Query().Get("item")
   price, ok := db[item]
   if !ok {
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprintf(w, "no such item: %q\n", item)
      return
   }
   fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
   item := req.URL.Query().Get("item")
   newprice := req.URL.Query().Get("price")
   _, ok := db[item]
   if !ok {
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprintf(w, "no such item: %q\n", item)
      return
   }
   newpricevalue, err := strconv.ParseFloat(newprice, 32)
   if err != nil {
      fmt.Fprintf(w, "invalid price: %q\n", newprice)
      return
   }
   db[item] = dollars(newpricevalue)
   fmt.Fprintf(w, "$%.2f\n", newpricevalue)
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
   item := req.URL.Query().Get("item")
   _, ok := db[item]
   if !ok {
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprintf(w, "no such item: %q\n", item)
      return
   }
   delete (db, item)
}

func main() {
   db := database{"shoes": 50, "socks": 5}
   http.HandleFunc("/list", db.list)
   http.HandleFunc("/create", db.create)
   http.HandleFunc("/read", db.read)
   http.HandleFunc("/update", db.update)
   http.HandleFunc("/delete", db.del)

   log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
