package main

import (
   "fmt"
   "log"
   "os"
   "time"
   "github"
)

func main() {
   result, err := github.SearchIssues(os.Args[1:])
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("%d issues: \n", result.TotalCount)
   var monthOld []*github.Issue
   var lessThanYearOld []*github.Issue
   var moreThanYearOld []*github.Issue

   currentYear := time.Now().Year()
   currentMonth := time.Now().Month()

   for _, item := range result.Items {
      year := item.CreatedAt.Year()
      month := item.CreatedAt.Month()

      if (year == currentYear && month == currentMonth) {
         monthOld = append(monthOld, item)
      } else if (year == currentYear) || (year == currentYear - 1 && month > currentMonth) {
         lessThanYearOld = append(lessThanYearOld, item)
      } else {
         moreThanYearOld = append(moreThanYearOld, item)
      }
   }

   fmt.Printf("Less than a month old:\n")
   for _, item := range monthOld {
      fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
   }

   fmt.Printf("Less than a year old:\n")
   for _, item := range lessThanYearOld {
      fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
   }

   fmt.Printf("More than a year old:\n")
   for _, item := range moreThanYearOld {
      fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
   }

}
