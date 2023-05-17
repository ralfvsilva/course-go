package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	fmt.Println("Valor de x:", x)

	// fmt.Println("Valor de x:" + x)  - identifica erro por considerar float64 (erro mismatched)

	xs := fmt.Sprint(x)
	fmt.Println("x:" + xs)

	fmt.Printf("x: %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v25", a, b, c, d)

}
