// Surface calcula uma renderização svg de uma função de superfície 3D 
package main

import (
   "fmt"
   "math"
   "io"
   "net/http"
   "strconv"
   "log"
   "os"
)

const (
   width, height = 600, 320 // tamanho do canvas em pixels
   cells         = 100      // número de células da grade
   xyrange       = 30.0     // intervalos dos eixos (-xyrange ... +xyrange)
   xyscale       = width / 2 / xyrange // pixels por unidade x ou y
   zscale        = height * 0.4 // pixels por unidade z
   angle         = math.Pi / 6  // ângulo dos eixos x, y (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30º), cosseno(30º)
var red, blue = "#ff0000", "#0000ff"
var color = blue

func main() {
   http.HandleFunc("/", handler)  // cada requisição chama handler
   log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, r *http.Request) {
   writer.Header().Set("Content-Type", "image/svg+xml")

   if err := r.ParseForm(); err != nil {
      log.Print(err)
   }

   var w, h int
   var c string
   for k, v := range r.Form {
      if k == "width" {
         w_, err := strconv.Atoi(v[0])
	 if err != nil {
	    fmt.Fprintf(os.Stdout, "error parsing parameter: %q\n", k)
	 } else {
	    w = w_
	 }
      } else if k == "height" {
         h_, err := strconv.Atoi(v[0])
	 if err != nil {
	    fmt.Fprintf(os.Stdout, "error parsing parameter: %q\n", k)
	 } else {
	    h = h_
	 }
      } else if k == "color" {
         if v[0] == "red" {
	    c = red
	 } else if v[0] == "blue" {
	    c = blue
	 }
      }
   }
   surface(writer, w, h, c)
}

func surface(out io.Writer, w int, h int, c string) {
   fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' " +
       "style='stroke: grey; fill: white; strokewidth: 0.7' " +
       "width='%d' height='%d'>", w, h)
   for i := 0; i < cells; i++ {
      for j := 0; j < cells; j++ {
         ax, ay := corner(i+1, j, w, h)
         bx, by := corner(i, j, w, h)
         cx, cy := corner(i, j+1, w, h)
         dx, dy := corner(i+1, j+1, w, h)
	 fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
	     ax, ay, bx, by, cx, cy, dx, dy, c)
      }
   }
   fmt.Fprintf(out, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
   // Encontra o ponto (x,y) no canto da célula(i,j)
   x := xyrange * (float64(i)/cells - 0.5)
   y := xyrange * (float64(j)/cells - 0.5)

   // Calcula a altura z da superfície
   z := f(x, y)

   // Faz uma projeção isométrica de (x, y, z) sobre (sx, sy) do canvas SVG 2D
   sx := float64(width)/2 + (x - y) * cos30 * xyscale
   sy := float64(height)/2 + (x + y) * sin30 * xyscale - z * zscale

   return sx, sy
}

func f(x, y float64) float64 {
   r := math.Hypot(x, y) // distância de (0, 0)
   return math.Sin(r) / r
}
