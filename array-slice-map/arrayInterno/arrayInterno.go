package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // slice de 10 posições e array interno com 20
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2) // Ambas as estruturas de slice apontam para o mesmo array
}
