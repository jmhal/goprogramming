package main

import (
   "fmt"
)

type tree struct {
   value string
   nodes []*tree
}

func main () {
   node7 := tree{value: "node7", nodes: []*tree{}}
   node6 := tree{value: "node6", nodes: []*tree{}}

   node5 := tree{value: "node5", nodes: []*tree{}}
   node4 := tree{value: "node4", nodes: []*tree{}}
   node3 := tree{value: "node3", nodes: []*tree{}}

   node1 := tree{value: "node1", nodes: []*tree{&node3, &node4, &node5}}
   node2 := tree{value: "node2", nodes: []*tree{&node6, &node7}}

   node0 := tree{value: "node0", nodes: []*tree{&node1, &node2}}

   fmt.Println("Breadth First:")
   breadthFirst(nodeInfo, []*tree{&node0})
   fmt.Println()
   fmt.Println("Depth First:")
   depthFirst(nodeInfo, &node0, make(map[*tree]bool))
   fmt.Println()

   return
}

func nodeInfo(node *tree) {
   fmt.Printf(" %s ", (*node).value)
}

func breadthFirst(f func(node *tree), worklist []*tree) {
   seen := make(map[*tree]bool)
   for len(worklist) > 0 {
      nodes := worklist
      worklist = nil
      for _, node := range nodes {
         if !seen[node] {
	    f(node)
	    seen[node] = true
	    worklist = append(worklist, (*node).nodes...)
	 }
      }
   }
}

func depthFirst(f func(node *tree), node *tree, visited map[*tree]bool) {
   visited[node] = true
   f(node)
   for _, child := range (*node). nodes {
      if !visited[child] {
         depthFirst(f, child, visited)
      }
   }
}
