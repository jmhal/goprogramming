package main

import (
  "fmt"
  "encoding/json"
  "log"
)

type NewIssue struct {
   Title     string
   Body      string
   Milestone int
   Assignees []string
   Labels     []string
}

func main() {
   var issue NewIssue
   issue.Title = "Teste"
   issue.Body = "Não sei qual o erro."
   issue.Milestone = 5
   issue.Assignees = []string{"jmhal"}
   issue.Labels = []string{"linux", "c++"}
   fmt.Printf("%s\n", issue)
   data, err := json.MarshalIndent(issue, "", "  ")
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("%s\n", data)
}
