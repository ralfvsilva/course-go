package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	// escrita
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou.")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // channel de inteiros e buffer com 3 posições
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch) // leitura
}
