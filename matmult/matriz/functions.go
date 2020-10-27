package matriz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matriz struct {
	naoLinear           [][]float64
	linear              []float64
	ordemPrincipalLinha bool
	ordem               int
}

func CriarMatriz(ordem int, ordemPrincipalLinha bool) Matriz {
	var matriz Matriz
	matriz.ordem = ordem
	matriz.ordemPrincipalLinha = ordemPrincipalLinha
	if ordemPrincipalLinha {
		matriz.linear = make([]float64, ordem*ordem)
		for i := 0; i < ordem*ordem; i++ {
			matriz.linear[i] = 1.0
		}

	} else {
		matriz.naoLinear = make([][]float64, ordem)
		for i := 0; i < ordem; i++ {
			matriz.naoLinear[i] = make([]float64, ordem)
			for j := 0; j < ordem; j++ {
				matriz.naoLinear[i][j] = 1.0
			}
		}
	}

	return matriz
}

func CriarArquivoMatriz(ordem int, matrizNome string) {
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

func ImprimeMatriz(matriz Matriz) {
	if matriz.ordemPrincipalLinha {
		for pos, elemento := range matriz.linear {
			if (pos > 0) && ((pos+1)%matriz.ordem == 0) {
				fmt.Printf("%.2f\n", elemento)
			} else {
				fmt.Printf("%.2f:", elemento)
			}
		}
	} else {
		for _, linha := range matriz.naoLinear {
			for _, elemento := range linha {
				fmt.Printf("%.2f ", elemento)
			}
			fmt.Println()
		}
	}
}

func LerMatriz(arquivo string, ordemPrincipalLinha bool) Matriz {
	var matriz Matriz
	matriz.ordemPrincipalLinha = ordemPrincipalLinha

	// Abrir arquivo da Matriz
	matrizArquivo, err := os.Open(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	// Ler a ordem da matriz
	scanner := bufio.NewScanner(matrizArquivo)
	scanner.Scan()
	ordem, err := strconv.Atoi(scanner.Text())

	// Aloca a matriz
	matriz.ordem = ordem
	if ordemPrincipalLinha {
		matriz.linear = make([]float64, ordem*ordem)
	} else {
		matriz.naoLinear = make([][]float64, ordem)
		for i := range matriz.naoLinear {
			matriz.naoLinear[i] = make([]float64, ordem)
		}
	}

	// Preenche a matriz
	linha := 0
	coluna := 0
	for scanner.Scan() {
		linhaCompleta := scanner.Text()
		for _, numero := range strings.Split(linhaCompleta, ":") {
			if ordemPrincipalLinha {
				matriz.linear[linha*ordem+coluna], err = strconv.ParseFloat(numero, 32)
			} else {
				matriz.naoLinear[linha][coluna], err = strconv.ParseFloat(numero, 32)
			}
			if err != nil {
				log.Fatal("Erro ao ler elemento da matriz.")
			}
			coluna++
		}
		coluna = 0
		linha++
	}

	return matriz
}

func GravarMatriz(matriz Matriz, matrizNome string) {
	arquivo, err := os.Create(matrizNome)
	if err != nil {
		log.Fatal("Impossível criar arquivo", err)
	}
	defer arquivo.Close()

	fmt.Fprintf(arquivo, "%d\n", matriz.ordem)

	if matriz.ordemPrincipalLinha {
		for pos, elemento := range matriz.linear {
			if (pos > 0) && ((pos+1)%matriz.ordem == 0) {
				fmt.Fprintf(arquivo, "%.2f\n", elemento)
			} else {
				fmt.Fprintf(arquivo, "%.2f:", elemento)
			}
		}
	} else {
		for _, linha := range matriz.naoLinear {
			for _, elemento := range linha {
				fmt.Fprintf(arquivo, "%.2f:", elemento)
			}
			if _, err := arquivo.Seek(-1, 1); err != nil {
				log.Fatal("Impossível retornar um caractere no arquivo", err)
			}
			fmt.Fprintf(arquivo, "\n")
		}
	}
}

func Matmult(matriz1 Matriz, matriz2 Matriz) Matriz {
	var matriz3 Matriz
	matriz3.ordemPrincipalLinha = matriz1.ordemPrincipalLinha
	matriz3.ordem = matriz1.ordem
	if matriz1.ordemPrincipalLinha && matriz2.ordemPrincipalLinha {
		ordem := matriz1.ordem
		matriz3.linear = make([]float64, ordem*ordem)
		for i := 0; i < ordem; i++ {
			for j := 0; j < ordem; j++ {
				for k := 0; k < ordem; k++ {
					matriz3.linear[i*ordem+j] += matriz1.linear[i*ordem+k] * matriz2.linear[k*ordem+j]
				}
			}
		}

	} else if (!matriz1.ordemPrincipalLinha) && (!matriz2.ordemPrincipalLinha) {
		ordem := matriz1.ordem
		matriz3.naoLinear = make([][]float64, ordem)
		for i := range matriz3.naoLinear {
			matriz3.naoLinear[i] = make([]float64, ordem)
		}

		for i := 0; i < ordem; i++ {
			for j := 0; j < ordem; j++ {
				for k := 0; k < ordem; k++ {
					matriz3.naoLinear[i][j] += matriz1.naoLinear[i][k] * matriz2.naoLinear[k][j]
				}
			}
		}

	} else {
		fmt.Println("Matrizes incompatíveis.")
	}
	return matriz3
}
