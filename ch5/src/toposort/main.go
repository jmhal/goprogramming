package main

import (
   "fmt"
   // "sort"
)

var prereqs = map[string][]string {
   "algorithms": {"data structures"},
   "linear algebra": {"calculus"},
   "compilers": {
      "data structures",
      "formal languages",
      "computer organization",
   },
   "data structures":      {"discrete math"},
   "databases":            {"data structures"},
   "discrete math":        {"intro to programming"},
   "formal languages":     {"discrete math"},
   "networks":             {"operating systems"},
   "operating systems":    {"data structures", "computer organization"},
   "computer organization":{"networks"},
   "programming languages":{"data structures", "computer organization"},
}

func main() {
   for i, course := range topoSort(prereqs) {
      fmt.Printf("%d:\t%s\n", i+1, course)
   }
}

func contains(s []string, e string) bool {
   for _, a := range s {
      if a == e {
         return true
      }
   }
   return false
}

func topoSort(m map[string][]string) map[int]string {
   var order int
   order = 0
   seen := make(map[string]bool)
   courses := make(map[int]string)
   var visitAll func(items []string)
   visitAll = func(items []string) {
      for _, item := range items {
         if !seen[item] {
	    seen[item] = true
	    visitAll(m[item])
	    courses[order] = item
	    order++
	 }
      }
   }

   var checkPreRequisites func(string, map[string][]string) []string
   var depth int
   var path []string
   checkPreRequisites = func(course string, m map[string][]string) []string {
      // fmt.Printf("%*s %s\n", depth * 2, "->",  course)
      path = append(path, course)
      depth++
      var dependencies []string
      for _, c := range m[course] {
	 if contains(path, c) {
	    fmt.Println("Circular dependency found!!!")
            fmt.Printf("Course: %s\n", path)
	    return []string{}
	 }
	 subdependencies := checkPreRequisites(c, m)
	 dependencies = append(dependencies, c)
	 dependencies = append(dependencies, subdependencies...)

      }
      i := 0
      for j, v := range path {
         if v == course {
	    i = j
	 }
      }
      path = append(path[:i], path[i+1:]...)
      depth--
      return dependencies
   }


   var keys []string
   for key := range m {
      keys = append(keys, key)
      checkPreRequisites(key, m)
   }

   // Precisa ser eliminado
   // sort.Strings(keys)
   visitAll(keys)
   return courses
}
