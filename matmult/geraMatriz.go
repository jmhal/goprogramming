package main

import (
   "fmt"
   "os"
   "log"
   "strconv"
)

func criarMatriz(ordem int, matrizNome string) {
   arquivo, err := os.Create(matrizNome)
   if err != nil {
      log.Fatal("Impossível criar arquivo", err)
   }
   defer arquivo.Close()

   fmt.Fprintf(arquivo, "%d\n", ordem)

   for i := 0; i < ordem; i++ {
      for j := 0; j < ordem; j++ {
         fmt.Fprintf(arquivo, "%.2f:", 1.0)
      }
      if _, err := arquivo.Seek(-1, 1); err != nil {
         log.Fatal("Impossível retornar um caractere no arquivo", err)
      }
      fmt.Fprintf(arquivo, "\n")
   }
}

func main() {
   matrizOrdem, err := strconv.Atoi(os.Args[1])
   if err != nil {
      log.Fatal("Impossível converter ordem", err)
   }
   matrizNome := os.Args[2]
   criarMatriz(matrizOrdem, matrizNome)
}


