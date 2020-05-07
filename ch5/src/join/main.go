package main

import (
   "fmt"
)

func join(sep string, strs ...string) string {
   if len(strs) == 0 {
      return ""
   }
   str := ""
   for _, s := range strs {
      str = str + sep + s
   }
   return str[1:]
}


func main() {
   fmt.Println(join(":", "casa", "comida", "e", "roupa", "lavada"))
}
