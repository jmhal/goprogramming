// Dup 2 exibe a contagem e o texto das linhas que aparecem mais de uma vez na entrada. Ele le de stdin ou de uma lista de arquivos nomeados.
package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

func main() {
   counts := make(map[string]int)
   belongs := make(map[string]string)
   files := os.Args[1:]
   if len(files) == 0 {
      countLines(os.Stdin, "entrada padrao", belongs, counts)
   } else {
      for _, arg := range files {
         f, err := os.Open(arg)
         if err != nil {
            fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
            continue
         }
         countLines(f, arg, belongs, counts)
         f.Close()
      }
   }
   
   for line, n := range counts {
      if n > 1 {
         fmt.Printf("%d\t%s\t:%s\n", n, line, belongs[line])
      }
   }
}

func countLines(f *os.File, filename string, belongs map[string]string, counts map[string]int){
   input := bufio.NewScanner(f)
   for input.Scan() {
      entry := input.Text()
      counts[entry]++
      
      files := strings.Split(belongs[entry], " ")
      lastfile := files[len(files)-1]
      if lastfile != filename {
         belongs[entry] += " " + filename
      }
   }
}
