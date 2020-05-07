package main

import (
   "fmt"
)

var value int

func noreturn(val int) {
   defer func() {
      if p := recover(); p != nil {
         value = 14
      }
   }()
   if val == 0 {
      panic("Problem.")
   }
}

func main() {
   noreturn(0)
   fmt.Println(value)
}
