// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
   "fmt"
   "os"
   "popcount"
   "strconv"
)

func main() {
   n, err := strconv.Atoi(os.Args[1])
   if err != nil {
      return
   }
   number := uint64(n)

   fmt.Printf("%d\n", popcount.PopCount(number))
   fmt.Printf("%d\n", popcount.PopCountLoop(number))
   fmt.Printf("%d\n", popcount.PopCount64(number))
   fmt.Printf("%d\n", popcount.PopCountClean(number))
}
