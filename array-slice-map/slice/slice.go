package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array - quando define o número do array é um array
	s1 := []int{1, 2, 3}  // slice - quando não há definição, é um slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! O Slice define um pedaço de um array.
	s2 := a2[1:3] // O Slice, pega os index de 1 a 3 só que não incluindo o 3, ou seja, [2, 3]
	fmt.Println(a2, s2)

	// O slice não gerou um outro array, e sim uma estrutura que aponta para estes numeros

	s3 := a2[:2] // Novo Slice, indo do início até o índice 2, não incluindo o index 2.
	fmt.Println(a2, s3)

	// É possível imaginar o slice como uma estrutura com um certo tamanho e terá um ponteiro -
	// que pega 1 elemento de um array e apartir desse elemento, ele pega uma estrutura -
	// contínua

	s4 := s2[:1]
	fmt.Println(s2, s4)
}
