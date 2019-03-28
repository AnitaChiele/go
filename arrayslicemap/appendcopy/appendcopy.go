package main

import (
	f "fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// append só pode ser usado pelos slices, a tentiva exemplificada dá um
	// erro de compilação: array1 = append(array1, 4, 5, 6) e também o primeiro
	// argumento deve ser um slice, então o exemplo ao lado também erro de compilação:
	// slice1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)

	f.Println(array1, slice1)

	slice2 := make([]int, 2)
	// copia o que tem no slice1 para o slice2 respeitado a quantidade definida no slice2,
	// ou seja, nesse exemplo, o slice2 vai ter 2 elementos e não 3 elementos como tem o
	// slice1. [slice2 = slice1].
	copy(slice2, slice1)

	// no copy só é possível utilizar um slice ou uma string para copiar, ou seja, o código
	// ao lado daria erro de compilação: copy(slice2, array1) | copy(array1, slice1)
	f.Println(slice1, slice2)
}
