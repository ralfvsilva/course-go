package main

import (
	"fmt"
	"html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagensx) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c

}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com.br"),
	)

	fmt.Println(<-c, " | ", <-c)
}
