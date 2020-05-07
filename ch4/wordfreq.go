// Counting word frequencies
package main

import (
   "fmt"
   "bufio"
   "os"
)

func main() {
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Split(bufio.ScanWords)
   freq := make(map[string]int) 
   for scanner.Scan() {
      freq[scanner.Text()]++
   }
   for k, v := range freq {
      fmt.Printf("%s %d\n", k, v)
   }
}
