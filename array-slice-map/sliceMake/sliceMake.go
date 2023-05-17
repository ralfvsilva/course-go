package main

import "fmt"

func main() {
	s := make([]int, 10) // Criado um slice de 10 posições, e internamente ele já cria o array associado
	s[9] = 12            // Adicionando o 12 na posição 10
	fmt.Println(s)

	s = make([]int, 10, 20) // Tamanho 10 no slice , mas a capacidade máxima do array é 20
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // tamanho do slice 21 e capacidade max do array 40
	fmt.Println(s, len(s), cap(s))
}
