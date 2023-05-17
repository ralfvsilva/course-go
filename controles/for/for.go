package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // nao tem while em Go
		fmt.Println()
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n Misturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Número par.")
		} else {
			fmt.Println("Número impar.")
		}
	}

	for {
		// laço infinito
		fmt.Println("Para sempre.")
		time.Sleep(time.Second)
	}
}
