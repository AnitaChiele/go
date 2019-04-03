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

	// outra forma de trabalhar com as funções anônimas são a declaração e a utilização direto:
	func(frase string) {
		f.Printf("Criando e usando funções anônimas: %s\n", frase)
	}("hello world!")

	// o parâmetro não é obrigatório
	func() {
		f.Printf("Criando e usando funções anônimas sem parâmetros!\n")
	}()
}
