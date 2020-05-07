// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
   "fmt"
   "os"
   "strconv"
   "tempconv"
   "lengthconv"
   "weightconv"
)

func main() {
   for _, arg := range os.Args[1:] {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
         fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	 os.Exit(1)
      }
      f := tempconv.Fahrenheit(t)
      c := tempconv.Celsius(t)
      
      m := lengthconv.Meters(t)
      ft := lengthconv.Feet(t)

      k := weightconv.Kilograms(t)
      p := weightconv.Pounds(t)

      fmt.Printf("Temperature: %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
      fmt.Printf("Length: %s = %s, %s = %s\n", m, lengthconv.MToF(m), ft, lengthconv.FToM(ft))
      fmt.Printf("Weight: %s = %s, %s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
   }
}
