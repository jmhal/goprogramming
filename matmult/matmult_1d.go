// Programa que lê e multiplica matrizes.
package main

import (
   "fmt"
   "os"
   "bufio"
   "log"
   "strconv"
   "strings"
//   "github.com/pkg/profile"
)

func lerMatriz(arquivo string) (int, []float64) {
   // Abrir arquivo da Matriz
   matrizArquivo, err := os.Open(arquivo)
   if err != nil {
      log.Fatal(err)
   }

   // Ler a ordem da matriz
   scanner := bufio.NewScanner(matrizArquivo)
   scanner.Scan()
   ordem, err := strconv.Atoi(scanner.Text())
   // fmt.Printf("Matriz de ordem: %d\n", ordem)

   // Declara a matriz
   matriz := make([]float64, ordem*ordem)

   // Preenche a matriz
   linha := 0
   coluna := 0
   for scanner.Scan() {
      linhaCompleta := scanner.Text()
      for _, numero := range strings.Split(linhaCompleta,":") {
         matriz[linha*ordem + coluna], err = strconv.ParseFloat(numero, 32)
	 coluna++
      }
      coluna = 0
      linha++
   }

  return ordem, matriz
}

func imprimeMatriz(ordem int, matriz []float64) {
   for pos, elemento := range matriz {
      if (pos > 0) && ((pos+1) % ordem == 0) {
         fmt.Printf("%.2f\n", elemento)
      } else {
         fmt.Printf("%.2f:", elemento)
      }
   }
}

func gravaMatriz(ordem int, matriz []float64, matrizNome string) {
   arquivo, err := os.Create(matrizNome)
   if err != nil {
      log.Fatal("Impossível criar arquivo", err)
   }
   defer arquivo.Close()

   fmt.Fprintf(arquivo, "%d\n", ordem)

   for pos, elemento := range matriz {
      if (pos > 0) && ((pos+1) % ordem == 0) {
         fmt.Fprintf(arquivo, "%.2f\n", elemento)
      } else {
         fmt.Fprintf(arquivo, "%.2f:", elemento)
      }
   }
}

func matmult(ordem int, matriz1 []float64, matriz2 []float64) []float64 {
   // Declara a matriz resultado
   matriz := make([]float64, ordem * ordem)
   for i := 0; i < ordem; i++ {
      for j := 0; j < ordem; j++ {
         for k := 0;  k < ordem; k++ {
            matriz[i*ordem + j] += matriz1[i*ordem + k] * matriz2[k*ordem + j]
         }
      }
   }
   return matriz
}


func main() {
   // CPU profiling by default
   // defer profile.Start(profile.MemProfile).Stop()

   // Descobre os nomes das matrizes.
   matriz1Nome := os.Args[1]
   matriz2Nome := os.Args[2]
   matriz3Nome := os.Args[3]
   fmt.Printf("Matriz 1: %s, Matriz 2: %s, Matriz 3: %s\n", matriz1Nome, matriz2Nome, matriz3Nome)

   ordem1, m1 := lerMatriz(matriz1Nome)
   ordem2, m2 := lerMatriz(matriz2Nome)
   if ordem1 != ordem2 {
      log.Fatal("Ordens diferentes!")
   }
   // imprimeMatriz(m1)
   // imprimeMatriz(m2)
   m3 := matmult(ordem1, m1, m2)
   // imprimeMatriz(m3)
   gravaMatriz(ordem1, m3, matriz3Nome)
}


