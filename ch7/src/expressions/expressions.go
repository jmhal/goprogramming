package main

import (
   "fmt"
   "eval"
   "bufio"
   "os"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
   fmt.Print("Enter expression:")
   expression, _ := reader.ReadString('\n')
   expr, err := eval.Parse(expression)
   fmt.Println("Parsing done.")
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println("Starting pretty print.")
   env := eval.Env{}
   fmt.Println(expr)
   expr.Environment(env)

   fmt.Println("Computing the Value.")
   got := fmt.Sprintf("%.6g", expr.Eval(env))
   fmt.Println(got)
}
