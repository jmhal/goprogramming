// Servidor de Imagens. 
package main

import (
   "log"
   "net/http"
   "image"
   "image/color"
   "image/gif"
   "io"
   "math"
   "math/rand"
   "time"
   "strconv"
   "os"
   "fmt"
)

var green = color.RGBA{0x1, 0x64, 0x1, 0x1}
var red = color.RGBA{0x64, 0x1, 0x1, 0x1}
var blue = color.RGBA{0x1, 0x1, 0x64, 0x1}
var black = color.Black
var palette = []color.Color{black, green, red, blue}

func main() {
   rand.Seed(time.Now().UTC().UnixNano())
   http.HandleFunc("/", handler)  // cada requisição chama handler
   log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
   if err := r.ParseForm(); err != nil {
      log.Print(err)
   }

   cycles := 5
   for k, v := range r.Form {
      if k == "cycles" {
         new_cycles, err := strconv.Atoi(v[0])
	 if err != nil {
	    fmt.Fprintf(os.Stdout, "error parsing parameter: %q\n", k)
	 }
	 cycles = new_cycles
      }
   }

   lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
   const (
      //cycles = 5   // número de revoluções completas do oscilador x
      res = 0.001  // resolução angular
      size = 200   // canvas da imagem cobre de [-size..+size]
      nframes = 64 // número de quadros da animação
      delay = 8    // tempo entre quadros em unidades de 10ms 
   )
   fmt.Fprintf(os.Stdout, "%d\n", cycles)
   freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
   anim := gif.GIF{LoopCount: nframes}
   phase := 0.0 // diferença de fase
   for i := 0; i < nframes ; i++ {
       rect := image.Rect(0, 0, 2*size + 1, 2*size + 1)
       img := image.NewPaletted(rect, palette)
       colorIndex :=  rand.Intn(3) + 1
       for t := 0.0; t < float64(cycles) * 2 * math.Pi; t += res {
          x := math.Sin(t)
          y := math.Sin(t * freq + phase)
	  img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), uint8(colorIndex))
       }
       phase += 0.1
       anim.Delay = append(anim.Delay, delay)
       anim.Image = append(anim.Image, img)
   }
   gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
