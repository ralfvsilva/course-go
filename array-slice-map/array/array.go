package main

import "fmt"

func main() {
	// homogêneo (mesmo tipo) e estático (tamanho fixo)

	var notas [3]float64

	notas[0], notas[1], notas[2] = 7.8, 5.5, 6.6

	fmt.Println(notas)

	var total float64 = 0
	for i := 0; i <= len(notas)-1; i++ {

		total += notas[i]
		fmt.Println(total)
	}

	fmt.Println("Total das notas:", total)

	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)

}
