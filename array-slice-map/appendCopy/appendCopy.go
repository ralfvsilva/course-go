package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) - não é possível, pois oo primeiro argumento deve swr
	// um slice e não um array

	fmt.Println(slice1)
	slice1 = append(slice1, 4, 5, 6, 7)
	fmt.Println(array1, slice1)

	s2 := array1[:2] // slice pegando até o index 2 do array1
	fmt.Println(s2)

	slice2 := make([]int, 2) // slice de inteiros com 2 elementos
	copy(slice2, slice1)     // slice2 recebe os elementos e o slice1 é a fonte
	fmt.Println(slice2)
}
