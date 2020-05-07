package main

import (
   "fmt"
)

func max(vals ...int) (int, error) {
   if len(vals) == 0 {
      return -1, fmt.Errorf("Cannot define the maximum of no numbers")
   }
   max := vals[0]
   for _, a := range vals {
      if a > max {
         max = a
      }
   }
   return max, nil
}

func min(vals ...int) (int, error) {
   if len(vals) == 0 {
      return -1, fmt.Errorf("Cannot define the minimum of no numbers")
   }
   min := vals[0]
   for _, a := range vals {
      if a < min {
         min = a
      }
   }
   return min, nil
}

func main() {
   maximum, err := max(3,2,5,1)
   if err == nil {
      fmt.Printf("Max(3 2 5 1): %d\n", maximum)
   }
   minimum, err := min(3,2,5,1)
   if err == nil {
      fmt.Printf("Min(3 2 5 1): %d\n", minimum)
   }
   new_maximum, err := max()
   if err != nil {
      fmt.Println(err)
      return
   } else {
      fmt.Println(new_maximum)
   }
}
