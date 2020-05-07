package main

import (
   "fmt"
   "os"
)

func main() {
   os.Chdir("/tmp")
   mydir := os.Getenv("PWD")
   fmt.Println(mydir)
}
