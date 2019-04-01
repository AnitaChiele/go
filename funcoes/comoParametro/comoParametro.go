package main

import (
	f "fmt"
)

func multiplicacao(a, b int) int {
	return a * b
}

// a função exec recebe uma função que tem dois parâmetros inteiros e
// retornará um int. A função exec também tem dois parâmetros inteiros:
// p1 e 02
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	f.Println(resultado)
}
