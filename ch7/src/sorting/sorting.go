package main

import (
   "time"
   "os"
   "text/tabwriter"
   "fmt"
   "sort"
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

func printTracks(tracks []*Track) {
   const format = "%v\t%v\t%v\t%v\t%v\t\n"
   tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
   fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
   fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
   for _, t := range tracks {
      fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
   }
   tw.Flush()
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

var tracks = []*Track {
   {"Go", "Delilah" , "From the Roots Up", 2012 , length("3m38s")},
   {"Go", "Moby", "Moby" , 1992, length("3m37s")},
   {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
   {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

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

func main() {
   /*
   fmt.Println("Sem ordenação:")
   printTracks(tracks)

   fmt.Println("TÍTULO:")
   sort.Sort(byTitle(tracks))
   printTracks(tracks)

   fmt.Println("ARTISTA:")
   sort.Sort(byArtist(tracks))
   printTracks(tracks)

   fmt.Println("ÁLBUM:")
   sort.Sort(byAlbum(tracks))
   printTracks(tracks)

   fmt.Println("ANO:")
   sort.Sort(byYear(tracks))
   printTracks(tracks)

   fmt.Println("DURAÇÃO:")
   sort.Sort(byLength(tracks))
   printTracks(tracks)
   */


   m := multiTierSort{tracks: tracks}
   sortClick(m, "Title")
   sortClick(m, "Artist")
   sortClick(m, "Album")
   sortClick(m, "Year")
   sortClick(m, "Length")
   printTracks(tracks)

   /*
   sort.Stable(byTitle(tracks))
   sort.Stable(byArtist(tracks))
   sort.Stable(byAlbum(tracks))
   sort.Stable(byYear(tracks))
   sort.Stable(byLength(tracks))
   printTracks(tracks)
   */
}
