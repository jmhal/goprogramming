package main

import (
  "os"
  "io"
  "fmt"
  "log"
  "strings"
  "strconv"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

var lastOne string

type Comic struct {
   Alt        string `json:"alt"`
   Title      string `json:"title"`
   SafeTitle  string `json:"safe_title"`
   Transcript string `json:"transcript"`
   Img        string `json:"img"`
}

func LoadComic(file string) (Comic) {
   var comic Comic
   comicFile, err := os.Open("comics/" + file)
   defer comicFile.Close()
   if err != nil {
      log.Fatal(err)
   }

   jsonParser := json.NewDecoder(comicFile)
   jsonParser.Decode(&comic)
   return comic
}

func ParseComics() ([]Comic) {
   var comics []Comic

   files, err := ioutil.ReadDir("comics/")
   if err != nil {
	log.Fatal(err)
   }
   for _, f := range files {
     comics = append(comics, LoadComic(f.Name()))
   }
   return comics
}

func DownloadFile(filepath string, url string) error {
   // Verify if file exists
   if _, err := os.Stat(filepath); err == nil {
      return nil
   }

   fmt.Printf("Retrieving comic %s...\n", filepath)
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

   url := "https://xkcd.com/"
   resp, err:= http.Get(url)
   if err != nil {
      log.Fatal(err)
   }
   bytes, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatal(err)
   }
   htmls := strings.Split(string(bytes), "\n")
   for _, v := range htmls {
      if strings.Contains(v, "Permanent link to this comic") {
         lastOne = strings.Split(v,"/")[3]
         fmt.Printf("%s\n", lastOne)
      }
   }

   last, err := strconv.Atoi(lastOne)
   if err != nil {
      log.Fatal(err)
   }
   for i := 1; i <= last; i++ {
      DownloadFile("comics/" + strconv.Itoa(i) + ".json", url + strconv.Itoa(i) + "/info.0.json")
   }

   resp.Body.Close()
   comics := ParseComics()
   for _, c := range comics {
      content := c.Title + c.SafeTitle + c.Transcript
      if strings.Contains(content, searchString) {
         fmt.Printf("URL: %s\n", c.Img)
         fmt.Printf("Transcript: %s\n", strings.Split(c.Transcript, "{")[0])
         fmt.Printf("\n")
      }
   }
}
