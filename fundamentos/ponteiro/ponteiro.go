package main

import "fmt"

func main() {
	i := 1

	// ponteiro é uma referência de memória

	var p *int = nil

	p = &i // pegando o endereço da variável i
	*p++   // desreferenciando (pegando o valor)
	i++

	// p++ - nao tem aritmetica de ponteiros

	fmt.Println(p, *p, i, &i)

	// *p tem acesso ao valor ao qual o ponteiro aponta
	// Apenas p tem acesso ao endereço que está armazenado dentro do ponteiro
}
