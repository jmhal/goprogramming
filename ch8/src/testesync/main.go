package main

import (
   "fmt"
   "sync"
)

func main() {
   var wg sync.WaitGroup

   wg.Add(1)
   go say(&wg, "Um")

   wg.Add(1)
   go say(&wg, "Dois")

   wg.Add(1)
   go say(&wg, "Três")

   wg.Wait()
   fmt.Println("Função Main")
   return
}

func say(wg *sync.WaitGroup, name string) {
   defer wg.Done()
   fmt.Println("Função " + name)
}
