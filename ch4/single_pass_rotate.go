package main

import (
   "fmt"
)

type Direction int
const (
   RIGHT Direction = iota
   LEFT
)

func rotate(s []int, d Direction) {
   if d == RIGHT {
      temp := s[len(s)-1]
      for i, j := len(s)-1, len(s)-2; i >= 1 && j >= 0; i , j = i-1, j-1 {
         s[i] = s[j]
      }
      s[0] = temp
   } else if d == LEFT {
      temp := s[0]
      for i, j := 0, 1; i <= len(s) - 2 && j <= len(s) - 1; i , j = i+1, j+1 {
         s[i] = s[j]
      }
      s[len(s) - 1] = temp
   }
}

func main() {
   vetor := [5]int {1, 2, 3, 4, 5}
   fatia := vetor[:]
   rotate(fatia, RIGHT)
   rotate(fatia, LEFT)
   fmt.Printf("%q\n", vetor)
}
