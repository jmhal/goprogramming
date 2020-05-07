package main

import (
   "fmt"
)

func main() {
   vetor := [5]int {1, 2, 3, 4, 5}
   fatia := vetor[:3]
   fmt.Printf("%q\n", vetor)
   fmt.Printf("%q\n", fatia)
   fmt.Printf("%d\n", vetor[3])
   fmt.Printf("%d\n", fatia[2])
   fmt.Printf("%d\n", len(fatia))
   fmt.Printf("%d\n", cap(fatia))
}
