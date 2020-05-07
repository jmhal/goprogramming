package main

import (
   "fmt"
   "bufio"
   "os"
   "strconv"
   "strings"
   "log"
   "bytes"
   "time"
   "net/http"
   "io/ioutil"
   "encoding/json"
   "os/exec"
)

type NewIssue struct {
   Title     string   `json:"title"`
   Body      string   `json:"body"`
   Milestone int      // `json:"milestone"`
   State     string   `json:"state"`
   Assignees []string `json:"assignees"`
   Labels    []string `json:"labels"`
}

type Issue struct {
   Number    int
   HTMLURL   string `json:"html_url"`
   Title     string
   State     string
   User      *User
   CreatedAt time.Time `json:"created_at"`
   Body      string
}

type User struct {
   Login   string
   HTMLURL string `json:"html_url"`
}

// API URL
const apiURL = "https://api.github.com"

// User and Repository Identification
var user string
var token string
var repo string

// Functions for dealing with issues
func createIssue(title string, body string, milestone int, assignees []string, labels []string) {
   fmt.Printf("Create Issue.\n")
   createURL := apiURL + "/repos/" + user + "/" + repo + "/issues"
   // fmt.Printf("Create URL: %s\n", createURL)
   issue := NewIssue{title, body, milestone, "open",  assignees, labels}
   // fmt.Printf("Issue: %s.\n", issue)
   data, err := json.Marshal(issue)
   if err != nil {
      log.Fatal(err)
   }
   // fmt.Printf("Data: %s.\n", data)

   req, err := http.NewRequest("POST", createURL, bytes.NewBuffer(data))
   if err != nil {
      log.Fatal(err)
   }
   req.Header.Set("Content-Type", "application/json")
   req.SetBasicAuth(user, token)

   client := &http.Client{}
   resp, err := client.Do(req)
   if err != nil {
      log.Fatal(err)
   }
   defer resp.Body.Close()

   // fmt.Println("response Status:", resp.Status)
   // fmt.Println("response Headers:", resp.Header)
   // bodyResp, _ := ioutil.ReadAll(resp.Body)
   // fmt.Println("response Body:", string(bodyResp))
}

func readIssue(issueId string) {
   fmt.Printf("Read Issue.\n")
   readURL := apiURL + "/repos/" + user + "/" + repo + "/issues/" + issueId
   resp, err := http.Get(readURL)
   if err != nil {
      log.Fatal(err)
   }

   if resp.StatusCode != http.StatusOK {
      resp.Body.Close()
      log.Fatal(resp.Status)
   }

   var result Issue
   if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
      resp.Body.Close()
      log.Fatal(err)
   }

   resp.Body.Close()

   fmt.Printf("---\n")
   fmt.Printf("Number: %d\n", result.Number)
   fmt.Printf("Title: %s\n", result.Title)
   fmt.Printf("State: %s\n", result.State)
   fmt.Printf("User: %s\n", result.User.Login)
   fmt.Printf("Body: %s\n", result.Body)
   fmt.Printf("---\n")
}

func updateIssue(issueId string, title string, body string, milestone int, state string,  assignees []string, labels []string) {
   fmt.Printf("Update Issue.\n")
   updateURL := apiURL + "/repos/" + user + "/" + repo + "/issues/" + issueId
   issue := NewIssue{title, body, milestone, state,  assignees, labels}
   // fmt.Printf("Issue: %s.\n", issue)
   data, err := json.Marshal(issue)
   if err != nil {
      log.Fatal(err)
   }
   // fmt.Printf("Data: %s.\n", data)

   req, err := http.NewRequest("POST", updateURL, bytes.NewBuffer(data))
   if err != nil {
      log.Fatal(err)
   }
   req.Header.Set("Content-Type", "application/json")
   req.SetBasicAuth(user, token)

   client := &http.Client{}
   resp, err := client.Do(req)
   if err != nil {
      log.Fatal(err)
   }
   defer resp.Body.Close()

   // fmt.Println("response Status:", resp.Status)
   // fmt.Println("response Headers:", resp.Header)
   // bodyResp, _ := ioutil.ReadAll(resp.Body)
   // fmt.Println("response Body:", string(bodyResp))
}

func deleteIssue(issueId string) {
   fmt.Printf("Delete Issue.\n")
   updateIssue(issueId, "closed", "", 1, "closed", []string{}, []string{})
}

func callEditor() (string) {
   cmd := exec.Command("/usr/bin/vim", "temp.txt")
   cmd.Stdout = os.Stdout
   cmd.Stdin = os.Stdin
   cmd.Stderr = os.Stderr
   cmd.Run()

   fileContent, err := ioutil.ReadFile("temp.txt")
   if err != nil {
      log.Fatal(err)
   }
   output := string(fileContent)

   err = os.Remove("temp.txt")
   if err != nil {
      log.Fatal(err)
   }

   return output
}

func main() {
   if (len(os.Args) != 4) {
      fmt.Printf("Usage: \n")
      fmt.Printf("./issues_manager user token repo\n")
      return
   }

   user = os.Args[1]
   repo = os.Args[2]
   token = os.Args[3]

   reader := bufio.NewReader(os.Stdin)
   var option int
   for ; option != 5; {
      fmt.Printf("Options:\n")
      fmt.Printf("1 - Create Issue.\n")
      fmt.Printf("2 - Read Issue.\n")
      fmt.Printf("3 - Update Issue.\n")
      fmt.Printf("4 - Delete Issue.\n")
      fmt.Printf("5 - Exit.\n")
      fmt.Printf("Choose One:")
      input, _ := reader.ReadString('\n')
      option, _ = strconv.Atoi(input[:len(input)-1])
      if option == 1 {
         fmt.Printf("Issue Title: ")
	 title, _ := reader.ReadString('\n')
         title = title[:len(title) - 1] 

	 body := callEditor()
	 fmt.Printf("\n")
	 fmt.Printf("Issue Assignees (comma separated): ")
	 assignees, _ := reader.ReadString('\n')
         assignees = assignees[:len(assignees)-1]
	 assigneesList := strings.Split(assignees, ",")

         fmt.Printf("Issue Labels (comma separated): ")
	 labels, _ := reader.ReadString('\n')
         labels = labels[:len(labels)-1]
	 labelsList := strings.Split(labels, ",")

         createIssue(title, body, 1, assigneesList , labelsList)
      } else if option == 2 {
         fmt.Printf("Issue: ")
	 issue, _ := reader.ReadString('\n')
         issue = issue[:len(issue)-1]
         readIssue(issue)
     } else if option == 3 {
         fmt.Printf("Issue: ")
	 issue, _ := reader.ReadString('\n')
         issue = issue[:len(issue)-1]

         fmt.Printf("Issue Title: ")
	 title, _ := reader.ReadString('\n')
         title = title[:len(title) - 1]

	 body := callEditor()
	 fmt.Printf("\n")
	 fmt.Printf("Issue Assignees (comma separated): ")
	 assignees, _ := reader.ReadString('\n')
         assignees = assignees[:len(assignees)-1]
	 assigneesList := strings.Split(assignees, ",")

         fmt.Printf("Issue Labels (comma separated): ")
	 labels, _ := reader.ReadString('\n')
         labels = labels[:len(labels)-1]
	 labelsList := strings.Split(labels, ",")

         updateIssue(issue, title, body, 1, "open", assigneesList, labelsList)
      } else if option == 4 {
         fmt.Printf("Issue: ")
	 issue, _ := reader.ReadString('\n')
         issue = issue[:len(issue)-1]

         deleteIssue(issue)
      }
   }
   fmt.Printf("Fim de execução.\n")
}
