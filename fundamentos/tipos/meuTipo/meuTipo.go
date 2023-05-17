package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool { // é preciso utilizar um type "nosso" como receiver, ou seja, como (n nota). (n float64) não pode.
	return float64(n) >= inicio && float64(n) <= fim
}

func notaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaConceito(9.5))
	fmt.Println(notaConceito(5.7))
	fmt.Println(notaConceito(7.9))
}
