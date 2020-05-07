package main

import (
   "fmt"
)

func reverse(s []int) {
   for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
      s[i], s[j] = s[j], s[i]
   }
}

func reverse_five(ptr *[5]int) {
   for i, j := 0, 4; i < j; i, j = i+1, j-1 {
      ptr[i], ptr[j] = ptr[j], ptr[i]
   }
}

func main() {
   vetor := [5]int {1, 2, 3, 4, 5}
   fatia := vetor[:3]
   reverse(fatia)
   reverse_five(&vetor)
   fmt.Printf("%q\n", vetor)
}
