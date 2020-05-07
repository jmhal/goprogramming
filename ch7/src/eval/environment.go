package eval

import (
   "bufio"
   "os"
   "fmt"
   "strconv"
   "strings"
)

func (v Var) Environment(env Env) Env {
   varname := string(v)
   reader := bufio.NewReader(os.Stdin)
   fmt.Printf("Enter value for %s :", varname)
   varinput, _ := reader.ReadString('\n')
   varinput = strings.TrimSuffix(varinput, "\n")
   varvalue, err := strconv.ParseFloat(varinput, 64)
   if err != nil {
      fmt.Println(err)
   }
   env[v] = varvalue
   return env
}

func (l literal) Environment(env Env) Env {
   return env
}

func (u unary) Environment(env Env) Env {
   return u.x.Environment(env)
}

func (b binary) Environment(env Env) Env {
   newEnv := Env{}
   for k, v := range b.x.Environment(env) {
      newEnv[k] = v
   }
   for k, v := range b.y.Environment(env) {
      newEnv[k] = v
   }
   return newEnv
}

func (c call) Environment(env Env) Env {
   newEnv := Env{}
   for _, a := range c.args {
      for k, v := range a.Environment(env) {
         newEnv[k] = v
      }
   }
   return newEnv
}


