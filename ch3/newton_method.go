package main

import (
   "fmt"
   "math/cmplx"
   "math"
)

func main() {
   z := complex(10,10)
   steps, root := newtonMethod(z)
   fmt.Printf("Number Of Steps: %d Root:%g Root:%g\n", steps, root, complex(round(real(root), 4), round(imag(root), 4)))
}

func newtonMethod(z complex128) (int, complex128) {
   const iterations = 37
   for i := 0; i < iterations; i++ {
      z -= (z - 1/(z*z*z)) / 4
      if cmplx.Abs(z*z*z*z-1) < 1e-6 {
         return i, z
      }
   }
   return 37, z
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
