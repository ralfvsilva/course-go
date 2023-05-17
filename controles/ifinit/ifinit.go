package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // pega o nano segundo atual
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // também suporta com o switch
		fmt.Println("Ganhou!")
		fmt.Println(numeroAleatorio())
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(numeroAleatorio())
	}
}
