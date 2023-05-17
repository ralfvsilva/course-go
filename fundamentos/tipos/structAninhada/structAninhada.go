package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item // slice porquê não está definido o tamanho nos colchetes
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10}, // exemplo mais simples para identificar o que está sendo passado
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
}
