package main

import "fmt"

func main() {
	// var aprovados map[int]string // chavo do map int e valor do tipo string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"
	aprovados[3] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // variavel cpf é a chave e nome o nome, vai percorrer aprovados
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1])
	delete(aprovados, 2)
	fmt.Println(aprovados[2])
	fmt.Println(aprovados[1])
}
