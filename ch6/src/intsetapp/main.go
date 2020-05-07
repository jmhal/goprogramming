package main

import (
   "fmt"
   "intset"
)

func main() {
   var x, y intset.IntSet
   x.Add(1)
   x.Add(144)
   x.Add(9)
   x.Add(10)
   fmt.Println(x.String())
   fmt.Println(x.Len())

   x.Add(9)
   y.Add(42)
   fmt.Println(y.String())
   fmt.Println(y.Len())

   x.UnionWith(&y)

   fmt.Println(x.String())
   fmt.Println(x.Len())
   z := x.Copy()

   fmt.Println(x.Has(9), x.Has(123))

   x.Remove(1)
   x.Remove(144)
   x.Add(1024)
   x.Add(10000)
   x.Remove(9)
   fmt.Println(x.String())
   x.Clear()

   fmt.Println(x.String())
   fmt.Println(z.String())
   x.AddAll(1,2,3,4,5)
   z.AddAll(1,2,3)
   fmt.Println(x.String())
   fmt.Println(z.String())

   s := z.SymmetricDifference(&x)
   fmt.Println(s.String())
   for _, e := range s.Elems() {
      fmt.Println(e)
   }
}
