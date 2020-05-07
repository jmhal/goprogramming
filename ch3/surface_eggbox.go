// Exercício 3.1
// Surface calcula uma renderização svg de uma função de superfície 3D
package main

import (
   "fmt"
   "math"
//   "os"
)

const (
   width, height = 600, 320 // tamanho do canvas em pixels
   cells         = 50      // número de células da grade
   xyrange       = 10.0     // intervalos dos eixos (-xyrange ... +xyrange)
   xyscale       = width / 2 / xyrange // pixels por unidade x ou y
   zscale        = height * 0.4 // pixels por unidade z
   angle         = math.Pi / 6  // ângulo dos eixos x, y (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30º), cosseno(30º)

func main() {
   fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
       "style='stroke: grey; fill: white; strokewidth: 0.7' " +
       "width='%d' height='%d'>", width*10, height*10)
   for i := 0; i < 5; i++ {
      for j := 0; j < 5; j++ {
         draw_egg(i, j)
      }
   }
   fmt.Println("</svg>")
}

func draw_egg(startx int, starty int) {
   for i := startx * cells; i < cells * (startx + 1); i++ {
      for j := starty * cells; j < cells * (starty + 1); j++ {
         ax, ay, err := corner(i+1, j)
   	 if !err {
   	    continue;
   	 }

         bx, by, err := corner(i, j)
   	 if !err {
   	    continue;
   	 }
   
         cx, cy, err := corner(i, j+1)
   	 if !err {
   	    continue;
   	 }
   
         dx, dy, err := corner(i+1, j+1)
   	 if !err {
   	    continue;
   	 }
   
   	 fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
   	    ax, ay, bx, by, cx, cy, dx, dy)
         }
   }
}

func corner(i, j int) (float64, float64, bool) {
   // Encontra o ponto (x,y) no canto da célula(i,j)
   x := xyrange * (float64(i)/cells - 0.5)
   y := xyrange * (float64(j)/cells - 0.5)

   // Calcula coordenadas auxiliares para a altura
   if i > 50 {
      i = i - (i / 50) * 50
   }
   if j > 50 {
      j = j - (j / 50) * 50
   }
   xa := xyrange * (float64(i)/cells - 0.5)
   ya := xyrange * (float64(j)/cells - 0.5)

   // Calcula a altura z da superfície
   z := f(xa, ya)

   // fmt.Fprintf(os.Stderr, "%f\n", z)

   // Verifica se o número gerado é finito ou não
   var isValid bool
   if math.IsInf(z, 0) {
      isValid = false
   } else {
      isValid = true
   }

   // Faz uma projeção isométrica de (x, y, z) sobre (sx, sy) do canvas SVG 2D
   sx := width/2 + (x - y) * cos30 * xyscale
   sy := height/2 + (x + y) * sin30 * xyscale - z * zscale

   return sx, sy, isValid
}

func f(x, y float64) float64 {
   r := math.Hypot(x, y) // distância de (0, 0)
   // fmt.Fprintf(os.Stderr, "%f\n", r)
   return math.Sin(r) / r
}
