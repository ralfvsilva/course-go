package main

import "fmt"

type produto struct { // tipo identificador(prooduto) struct
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	// Outra forma de declarar struct

	produto2 := produto{"Notebook", 1899.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
