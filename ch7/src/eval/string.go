package eval

import (
   "fmt"
)

var depth int

func (v Var) String() string {
   depth++
   output := fmt.Sprintf("depth:%d %*s%s\n", depth,  depth, "",  string(v))
   depth--
   return output
}

func (l literal) String() string {
   depth++
   output := fmt.Sprintf("depth:%d %*s%0.2f\n", depth, depth, "",  l)
   depth--
   return output
}

func (u unary) String() string {
   depth++
   output := fmt.Sprintf("depth:%d %*s%s\n%s", depth, depth, "", string(u.op), u.x)
   depth--
   return output
}

func (b binary) String() string {
   depth++
   output := fmt.Sprintf("depth:%d %*s%v\n%s%s", depth, depth, "", string(b.op), b.x, b.y)
   depth--
   return output
}

func (c call) String() string {
   depth++
   output := fmt.Sprintf("depth:%d %*s%s\n", depth, depth, "", c.fn)
   for _, a := range c.args {
      output += fmt.Sprintf("%s", a)
   }
   depth--
   return output
}


