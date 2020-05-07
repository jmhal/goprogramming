package main

import (
  "os"
  "io"
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type Movie struct {
   Title   string
   Poster  string
}

const apiURL = "http://www.omdbapi.com/?apikey=6b2d0aa7&"

func DownloadFile(filepath string, url string) error {
   // Verify if file exists
   if _, err := os.Stat(filepath); err == nil {
      return nil
   }

   // Create the file
   out, err := os.Create(filepath)
   if err != nil {
       return err
   }
   defer out.Close()

   // Get the data
   resp, err := http.Get(url)
   if err != nil {
       return err
   }
   defer resp.Body.Close()

   // Write the body to file
   _, err = io.Copy(out, resp.Body)
   if err != nil {
       return err
   }

   return nil
}


func main() {
   searchString := os.Args[1]
   readURL := apiURL + "t=" + searchString
   resp, err := http.Get(readURL)
   if err != nil {
      log.Fatal(err)
   }

   if resp.StatusCode != http.StatusOK {
      resp.Body.Close()
      log.Fatal(resp.Status)
   }

   var result Movie
   if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
      resp.Body.Close()
      log.Fatal(err)
   }
   resp.Body.Close()
   fmt.Printf("%s\n", result.Poster)
   err = DownloadFile("movie.jpg", result.Poster)
   if err != nil {
      log.Fatal("No poster found.")
   }
}
