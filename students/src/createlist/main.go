package main

import (
   "fmt"
   "os"
   "golang.org/x/net/html"
   "strings"
   "regexp"
   "sort"
)

type Student struct {
   id    string
   name  string
   email string
}

type ByName []Student

func (s ByName) Len() int           { return len(s) }
func (s ByName) Less(i, j int) bool { return s[i].name < s[j].name }
func (s ByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
   doc, err := html.Parse(os.Stdin)
   if err != nil {
      fmt.Fprintf(os.Stderr, "createlist: %v\n", err)
      os.Exit(1)
   }
   var elements []Student
   elements = visit(elements, doc)

   sort.Sort(ByName(elements))

   for _,s := range elements {
      fmt.Printf("%s;%s;%s\n", s.id, s.name, s.email)
   }

}

func visit(elements []Student, n *html.Node) []Student {
   if n != nil {
      if n.Type == html.ElementNode {
         for _, a := range n.Attr {
            if a.Key == "valign"{
               var student Student
               for c := n.FirstChild; c != nil; c = c.NextSibling {
                  if c.Data == "strong"  {
                     for c1 := c.FirstChild; c1 != nil; c1 = c1.NextSibling {
                        if c1.Data == "a" || c1.Data == "img" {
                           continue
                        } else {
                           name := strings.TrimSpace(c1.Data)
                           number, _ := regexp.MatchString(`[0-9]+`, name)
                           if name != "" && !number {
                              student.name = name
                           }
                        }
                     }
                  } else if c.Data == "em" {
                     for c1 := c.FirstChild; c1 != nil; c1 = c1.NextSibling {
                        if c1.Data == "a" || c1.Data == "img" {
                           continue
                        } else {
                           data := strings.TrimSpace(c1.Data)
                           id, _ := regexp.MatchString(`^[0-9]+$`, data)
                           email, _ := regexp.MatchString(`@`, data)
                           if data != "" {
                              if id {
                                 student.id = data
                              }
                              if email {
                                 student.email = data
                              }
                           }
                        }
                     }
                  }
               }
               if student.name != "" && student.id != "" && student.email != "" {
                  elements = append(elements, student)
               }
            }
         }
      }
      if n.NextSibling != nil {
         elements = visit(elements, n.NextSibling)
      }
      if n.FirstChild != nil {
         elements = visit(elements, n.FirstChild)
      }
   }
   return elements
}
