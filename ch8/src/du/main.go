package main

import (
   "flag"
   "fmt"
   "io/ioutil"
   "os"
   "sync"
   "time"
   "path/filepath"
)

// Semaphore to control the number of open files
var sema = make(chan struct{}, 20)

func main() {
   // Determine the initial directories.
   flag.Parse()
   roots := flag.Args()
   if len(roots) == 0 {
      roots = []string{"."}
   }

   period := time.Tick(2 * time.Second)
   for {
      select {
      case <- period:
         // Traverse each root of the file tree in parallel
         fileSizes := make([]chan int64, len(roots))
         for i := 0 ; i < len(roots); i++ {
            fileSizes[i] = make(chan int64)
         }
         var n sync.WaitGroup
         for index, root := range roots {
            n.Add(1)
            go walkDir(root, &n, fileSizes[index])
         }
         go func() {
            n.Wait()
            for index, _ := range roots {
               close(fileSizes[index])
            }
         }()

         var p sync.WaitGroup
         for index, root := range roots {
            p.Add(1)
            go printPartials(root, fileSizes[index], &p)
         }
         p.Wait()
         fmt.Println()
      }
   }
}

func printPartials(root string, fileSizes <-chan int64, p *sync.WaitGroup){
   defer p.Done()
   var tick <-chan time.Time
   tick = time.Tick(500 * time.Millisecond)
   var nfiles, nbytes int64

loop:
   for {
      select {
      case size, ok := <-fileSizes:
         if !ok {
            break loop // fileSizes was closed
         }
         nfiles++
         nbytes += size
      case <- tick:
         // printDiskUsage(root, nfiles, nbytes)
      }
   }
   printDiskUsage(root, nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
   defer n.Done()
   for _, entry := range dirents(dir) {
      if entry.IsDir() {
         n.Add(1)
         subdir := filepath.Join(dir, entry.Name())
         go walkDir(subdir, n, fileSizes)
      } else {
         fileSizes <- entry.Size()
      }
   }
}

func dirents(dir string) []os.FileInfo {
   sema <- struct{}{}
   defer func(){ <-sema } ()
   entries, err := ioutil.ReadDir(dir)
   if err != nil {
      fmt.Fprintf(os.Stderr, "du: %v\n", err)
      return nil
   }
   return entries
}

func printDiskUsage(root string, nfiles, nbytes int64) {
   fmt.Printf("%s: %d files %.1f MB\n", root, nfiles, float64(nbytes)/1048576)
}
