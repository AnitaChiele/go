package main

import (
	f "fmt"
)

func main() {
	// homogênia (mesmo tipo) e estática (fixo em tamanho)
	// não é necessário inicializá-lo, já vem com os valores
	// "zerados" de acordo com o tipo de dado do array.
	var notas [3]float64

	f.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	f.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	f.Printf("Media: %.2f", media)
}
