package main

import "fmt"

func imprimirResultado(notaFinal float64) {

	// nao coloca parenteses na estrutura
	if notaFinal >= 7 {
		fmt.Println("Aprovado com nota", notaFinal)
	} else {
		fmt.Println("Reprovado, não atingiu a nota necessária,", notaFinal)
	}
}

func main() {
	imprimirResultado(6.9)
}
