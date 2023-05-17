package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"Jose joao":      11300.50,
		"Gabriela Silva": 15432.20,
		"Pedro Junior":   1200.0,
	}

	funcESalarios["Rafael Filho"] = 1350.59
	delete(funcESalarios, "Inexistente") // Não gera erro, ele tenta excluir

	for nome, salario := range funcESalarios {
		fmt.Printf("%s: (Salário) R$ %.2f\n", nome, salario)
	}

}
