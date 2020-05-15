// Programa que lê e multiplica matrizes.
package main

import (
   "fmt"
   "os"
   "bufio"
   "log"
   "strconv"
   "strings"
)

func lerMatriz(arquivo string) [][] float64 {
   // Abrir arquivo da Matriz
   matrizArquivo, err := os.Open(arquivo)
   if err != nil {
      log.Fatal(err)
   }

   // Ler a ordem da matriz
   scanner := bufio.NewScanner(matrizArquivo)
   scanner.Scan()
   ordem, err := strconv.Atoi(scanner.Text())
   fmt.Printf("Matriz de ordem: %d\n", ordem)
   
   // Declara a matriz
   matriz := make([][]float64, ordem)
   for i := range matriz {
      matriz[i] = make([]float64, ordem)
   }

   // Preenche a matriz
   linha := 0
   coluna := 0
   for scanner.Scan() {
      linhaCompleta := scanner.Text()
      for _, numero := range strings.Split(linhaCompleta,":") {
         matriz[linha][coluna], err = strconv.ParseFloat(numero, 32)
	 coluna++
      }
      coluna = 0 
      linha++
   }

  return matriz
}

func imprimeMatriz(matriz [][]float64) {
   for _, linha := range matriz {
      for _, elemento := range linha {
         fmt.Printf("%.2f ", elemento)
      }
      fmt.Println()
   }
}

func matmult(matriz1 [][]float64, matriz2 [][]float64) [][]float64 {
   // Recupera a ordem 
   ordem := len(matriz1)

   // Declara a matriz resultado
   matriz := make([][]float64, ordem)
   for i := range matriz {
      matriz[i] = make([]float64, ordem)
   }

   for i := 0; i < ordem; i++ {
      for j := 0; j < ordem; j++ {
         for k := 0;  k < ordem; k++ {
	    matriz[i][j] += matriz1[i][k] * matriz2[k][j]
	 }
      }
   }
   return matriz
}

func main() {
   // Descobre os nomes das matrizes.
   matriz1Nome := os.Args[1]
   matriz2Nome := os.Args[2]
   matriz3Nome := os.Args[3]
   fmt.Printf("Matriz 1: %s, Matriz 2: %s, Matriz 3: %s\n", matriz1Nome, matriz2Nome, matriz3Nome)

   m1 := lerMatriz(matriz1Nome)
   m2 := lerMatriz(matriz2Nome)
   imprimeMatriz(m1)
   imprimeMatriz(m2)
   m3 := matmult(m1, m2)
   imprimeMatriz(m3)
}


