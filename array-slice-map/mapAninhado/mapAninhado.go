package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{

		"G": {
			"Gabriela Silva": 1543.40,
			"Guga Pereira":   1600.30,
		},
		"J": {
			"José João": 1430.00,
		},
		"P": {
			"Pedro Junior": 1350.40,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
