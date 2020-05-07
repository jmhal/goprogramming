package main

import (
   "time"
   "sort"
   "net/http"
   "html/template"
   "log"
)

type Track struct {
   Title string
   Artist string
   Album string
   Year int
   Length time.Duration
}

type byTitle  []*Track
type byArtist []*Track
type byAlbum  []*Track
type byYear   []*Track
type byLength []*Track

func length(s string) time.Duration {
   d, err := time.ParseDuration(s)
   if err != nil {
      panic(s)
   }
   return d
}

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type multiTierSort struct {
   tracks []*Track
   tiers []string
}

func (x multiTierSort) Len() int { return len(x.tracks) }
func (x multiTierSort) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }
func (x multiTierSort) Less(i, j int) bool {
   for _, tier := range x.tiers {
      switch tier {
         case "Title" :
	    if x.tracks[i].Title == x.tracks[j].Title {
	       continue;
	    } else {
	       return x.tracks[i].Title < x.tracks[j].Title
	    }
            case "Artist" :
            if x.tracks[i].Artist == x.tracks[j].Artist {
	       continue;
	    } else {
	       return x.tracks[i].Artist < x.tracks[j].Artist
	    }
	 case "Album" :
            if x.tracks[i].Album == x.tracks[j].Album {
	       continue;
	    } else {
	       return x.tracks[i].Album < x.tracks[j].Album
	    }
	 case "Year" :
	    if x.tracks[i].Year == x.tracks[j].Year {
	       continue;
	    } else {
	       return x.tracks[i].Year < x.tracks[j].Year
	    }
	 case "Length" :
	    if x.tracks[i].Length == x.tracks[j].Length {
	       continue;
	    } else {
	       return x.tracks[i].Length < x.tracks[j].Length
	    }
	 case "" :
	    break
      }
   }
   return true
}

func sortClick(m multiTierSort, key string) {
   if len(m.tiers) == 0 {
      m.tiers = make([]string, 5)
   }
   m.tiers = m.tiers[1:]
   m.tiers[0] = key
   sort.Sort(m)
}

var tracks = []*Track {
   {"Go", "Delilah" , "From the Roots Up", 2012 , length("3m38s")},
   {"Go", "Moby", "Moby" , 1992, length("3m37s")},
   {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
   {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var multisort = multiTierSort{tracks: tracks}

const index_html = `<html>
   <head>
      <style>
      table, th, td {
         border: 1px solid black;
      }
      </style>
      <title>
         Sorting Songs
      </title>
   </head>
   <body>
      <table>
         <tr>
	    <th> <a href="/Title">Title</a> </th>
	    <th> <a href="/Artist">Artist</a> </th>
	    <th> <a href="/Album"> Album </a> </th>
	    <th> <a href="/Year">Year</a> </th>
	    <th> <a href="/Length">Length</a> </th>
	 </tr>
	 {{ range .Tracks }}
	 <tr> 
	    <td> {{.Title}} </td>
	    <td> {{.Artist}} </td>
	    <td> {{.Album}} </td>
	    <td> {{.Year}} </td>
	    <td> {{.Length}} </td>
	 </tr>
	 {{end}}
      </table>
   </body>
</html>
`

func printTracks() {
   // Redirecionar a raiz para página que carrega o template do index
   http.HandleFunc("/", index)

   http.HandleFunc("/Title", getSortFunc("Title"))
   http.HandleFunc("/Artist", getSortFunc("Artist"))
   http.HandleFunc("/Album", getSortFunc("Album"))
   http.HandleFunc("/Year", getSortFunc("Year"))
   http.HandleFunc("/Length", getSortFunc("Length"))

   // Ativa o servidor
   log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func convert(original []*Track) ([]Track) {
   tracks := make([]Track, len(original))
   for i, t := range original {
      tracks[i] = Track{(*t).Title, (*t).Artist, (*t).Album, (*t).Year, (*t).Length}
   }
   return tracks
}

func getSortFunc(funcName string) func(http.ResponseWriter, *http.Request) {
   return func (w http.ResponseWriter, r *http.Request) {
      sortClick(multisort, funcName)
      orderedTracks := convert(tracks)
      data := struct {
         Tracks []Track
      }{
         orderedTracks,
      }

      tmpl := template.Must(template.New("index").Parse(index_html))
      tmpl.Execute(w, data)
   }
}

func title(w http.ResponseWriter, r *http.Request) {
   sortClick(multisort, "Title")
   orderedTracks := convert(tracks)
   data := struct {
      Tracks []Track
   }{
      orderedTracks,
   }

   tmpl := template.Must(template.New("index").Parse(index_html))
   tmpl.Execute(w, data)
}

func artist(w http.ResponseWriter, r *http.Request) {
   sortClick(multisort, "Artist")
   orderedTracks := convert(tracks)
   data := struct {
      Tracks []Track
   }{
      orderedTracks,
   }

   tmpl := template.Must(template.New("index").Parse(index_html))
   tmpl.Execute(w, data)
}


func index(w http.ResponseWriter, r *http.Request) {
   orderedTracks := convert(tracks)
   data := struct {
      Tracks []Track
   }{
      orderedTracks,
   }

   tmpl := template.Must(template.New("index").Parse(index_html))
   tmpl.Execute(w, data)
}



func main() {
   /*
   m := multiTierSort{tracks: tracks}
   sortClick(m, "Title")
   sortClick(m, "Artist")
   sortClick(m, "Album")
   sortClick(m, "Year")
   sortClick(m, "Length")
   */
   printTracks()
}
