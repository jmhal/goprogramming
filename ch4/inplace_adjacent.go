package main

import (
   "fmt"
)

type Direction int
const (
   RIGHT Direction = iota
   LEFT
)

func rotate(s []string, d Direction) {
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

func adjacent(strs []string) int {
   var match int;
   for i := 0; i < len(strs) - match; i++ {
      for strs[i] == strs[i+1] {
         rotate(strs[i:], LEFT)
	 match++
      }
   }
   return match
}

func main() {
   vetorString := [10]string { "aaa", "aaa", "aaa", "ddd", "bbb", "ccc", "ccc", "eee", "fff", "fff" }
   match := adjacent(vetorString[:])
   fmt.Printf("%q\n", vetorString)
   fmt.Printf("%q\n", vetorString[:len(vetorString) - match])
}
