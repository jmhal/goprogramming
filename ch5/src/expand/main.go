package main

import (
   "fmt"
   "os"
   "strings"
)

func main() {
   str := os.Args[1]
   fmt.Println(expand(str, joker))
}

func expand(s string, f func(string)string) string {
   arguments := strings.Split(s, "$")
   var results []string
   for i, a := range arguments {
      if (i == 0) {
         results = append(results, a)
      } else {
         results = append(results, f(a))
      }
   }
   return strings.Join(results, "")
}

func joker(s string) string {
   return "HaHAhAhAhaHA" + s
}

