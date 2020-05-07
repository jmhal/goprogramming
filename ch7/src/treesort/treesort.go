package main

import (
   "fmt"
   "strconv"
)

type tree struct {
   value int
   left, right *tree
}

func (t *tree) String() string {
   return getValues(t)
}

func getValues(t *tree) (string) {
   output := ""
   output += ":" + strconv.Itoa(t.value)
   if t.left != nil {
      output += ":" + getValues(t.left)
   }
   if t.right != nil {
      output += ":" + getValues(t.right)
   }
   return output
}

func Sort(values []int) {
   var root *tree
   for _, v := range values {
      root = add(root, v)
   }
   appendValues(values[:0], root)
}

func appendValues(values []int, t* tree) []int {
   if t != nil {
      values = appendValues(values, t.left)
      values = append(values, t. value)
      values = appendValues(values, t.right)
   }

   return values
}

func add(t *tree, value int) *tree {
   if t == nil {
      t = new(tree)
      t.value = value
      return t
   }
   if value < t.value {
      t.left = add(t.left, value)
   } else {
      t.right = add(t.right, value)
   }
   return t
}

func main() {
   var arv tree
   add(&arv, 1)
   add(&arv, 2)
   add(&arv, 100)
   add(&arv, 3)
   fmt.Println(&arv)

}
