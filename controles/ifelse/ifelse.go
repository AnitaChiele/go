package main

import (
	f "fmt"
)

func imprimirResultado(nota float64) {
	// pode utilizar parenteses apenas para forçar
	// precedência, caso contrário não é uma boa prática
	// colocar parenteses por fora do if - que envolve todo
	// if.

	if nota >= 7 {
		f.Println("Aprovado com nota:", nota)
	} else {
		f.Println("Reprovado com nota", nota)
	}

	if nota >= 7 && (true || false) {
		f.Println("Assim é okay.")
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
