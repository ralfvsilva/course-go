package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal (escrita)

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 5 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // assincronismo
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b) // sincronismo

	fmt.Println(<-c)
	fmt.Println(<-c)
}
