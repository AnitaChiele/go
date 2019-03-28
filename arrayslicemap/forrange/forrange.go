package main

import (
	f "fmt"
)

func main() {
	// os [...] no início do array diz pro compilador
	// contar quantos elementos o array possui e fixar
	// como lenght elementos - 1.
	numeros := [...]int{5, 6, 7, 8, 9}

	for indice, numero := range numeros {
		f.Printf("%d) %d\n", indice, numero)
	}

	println("----------------------")

	for i := range numeros {
		f.Printf("%d\n", i)
	}

	println("----------------------")

	for _, numero := range numeros {
		f.Printf("%d\n", numero)
	}
}
