package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Os ... deixa a definição do index do array ->
	// com o número de valores que são passados nas chaves

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros { // forRange sem var i, substituido por _, num é o ->
		// apelido que recebe o range do array numeros
		fmt.Println(num)
	}
}
