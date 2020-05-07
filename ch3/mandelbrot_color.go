// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// Color supported version
package main
import (
   "image"
   "image/color"
   "image/png"
   "math/cmplx"
   "os"
)

var turn uint8

func main() {
   const (
      xmin, ymin, xmax, ymax = -2, -2, +2, +2
      width, height = 1024, 1024
   )
   turn = 0
   img := image.NewRGBA(image.Rect(0, 0, width, height))
   for py := 0; py < height; py++ {
      y := float64(py)/height*(ymax-ymin) + ymin
      for px := 0; px < width; px++ {
         x := float64(px)/width * (xmax - xmin) + xmin
	 z := complex(x, y)
	 // Image point (px, py) represents complex value z
	 img.Set(px, py, mandelbrot(z))
      }
   }
   png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
   const iterations = 200
   const contrast = 15

   var v complex128
   for n := uint8(0); n < iterations; n++ {
      v = v * v + z
      if cmplx.Abs(v) > 2 {
         if n % 2 == 0 {
	    return color.RGBA{R:(255 - n), G:n, B:n, A:(contrast * n)}
	 } else if n % 3 == 0 {
	    return color.RGBA{R:n, G:(255 - n), B:n, A:(contrast * n)}
	 } else {
	    return color.RGBA{R:n, G:n, B:(255 - n), A:(contrast * n)}
	 }
      }
   }
   turn++
   return color.RGBA{R:turn, G:turn+turn, B:turn*turn, A:255}
}
