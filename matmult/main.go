package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmhal/goprogramming/matmult/matriz"
)

func main() {
	ordem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Ordem inválida.")
		return
	}
	fmt.Println("Define duas matrizes lineares.")
	m1 := matriz.CriarMatriz(ordem, true)
	m2 := matriz.CriarMatriz(ordem, true)

	fmt.Println("Multiplica as matrizes lineares.")
	start := time.Now()
	matriz.Matmult(m1, m2)
	end := time.Since(start)
	fmt.Printf("Duração: %s\n", end)

	fmt.Println("Define duas matrizes não lineares.")
	m1 = matriz.CriarMatriz(ordem, false)
	m2 = matriz.CriarMatriz(ordem, false)

	fmt.Println("Multiplica as matrizes não lineares.")
	start = time.Now()
	matriz.Matmult(m1, m2)
	end = time.Since(start)
	fmt.Printf("Duração: %s\n", end)
}
