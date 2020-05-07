package main

import (
   "fmt"
   "os"
   "io"
   "encoding/xml"
   "strings"
)

type Node interface{}

type CharData string

type Element struct {
   Type xml.Name
   Attr []xml.Attr
   Children []Node
}

func (e *Element) String() string {
   return fmt.Sprintf("%s %s %s,", e.Type, e.Attr, e.Children)
}

var depth int

func printTree(node Node) {
   depth++
   switch node := node.(type) {
      case Element:
         fmt.Printf("%d %*s\n", depth,  depth*5, node.Type.Local)
         for _, n := range node.Children{
            printTree(n)
         }
      case CharData:
            fmt.Printf("%d %*s\n", depth, depth*5, node)
   }
   depth--
}

func main () {
   dec := xml.NewDecoder(os.Stdin)
   var root *Node
   var stack []Node
   for {
      tok, err := dec.Token()
      if err == io.EOF {
         break
      } else if err != nil {
         fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
	 os.Exit(1)
      }
      switch tok := tok.(type) {
         case xml.StartElement :
	    // fmt.Printf("StartElement: %s\n", tok.Name.Local)
	    stack = append(stack, &Element{tok.Name, tok.Attr, []Node{}})
	    // fmt.Printf("Top of the Stack: %s\n", stack[len(stack)-1])
	    // fmt.Printf("Stack: %s\n", stack)
	 case xml.EndElement :
	    //goon = false
	    // fmt.Printf("EndElement: %s\n", tok.Name.Local)
	    if len(stack) < 2 {
	       root = &stack[0]
	       break
	    }
	    // fmt.Printf("CurrentNode: %s\n", stack[len(stack)-1].(*Element))
	    // fmt.Printf("ParentNode: %s\n", stack[len(stack)-2].(*Element))
	    parent := stack[len(stack)-2].(*Element)
	    current := stack[len(stack)-1].(*Element)
	    (*parent).Children = append((*parent).Children, *current)
	    // fmt.Printf("ParentNode: %s\n", *parent)
	    stack = stack[:len(stack)-1]
	    // fmt.Printf("Stack: %s\n", stack)
	 case xml.CharData:
	    if len(strings.TrimSpace(string(tok))) == 0 {
	       continue
	    }
	    // fmt.Printf("CharData: %s (%d)\n", tok, len(tok))
	    parent := stack[len(stack)-1].(*Element)
	    // fmt.Printf("Parent before CharData: %s\n", parent)
	    parent.Children = append(parent.Children, CharData(strings.TrimSpace(string(tok))))
	    // fmt.Printf("Parent after CharData: %s\n", parent)
      }
   }
   fmt.Println("ROOT:")
   fmt.Println((*root).(*Element).Type)
   for _, c := range (*root).(*Element).Children {
     printTree(c)
   }
}
