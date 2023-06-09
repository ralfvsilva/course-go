package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool { // é preciso utilizar um type "nosso" como receiver, ou seja, como (n nota). (n float64) não pode.
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(5.7))
	fmt.Println(notaParaConceito(7.9))
}
