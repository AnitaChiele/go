package main

import (
	f "fmt"
)

/*
	Função anônima não possuem nomes, porém é possível chamá-las através
	da variável que armazena ela, no caso baixo: soma.
*/
var soma = func(a, b int) int {
	return a + b
}

func main() {
	f.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	f.Println(sub(2, 3))
}
