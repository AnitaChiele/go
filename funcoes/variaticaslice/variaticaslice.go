package main

import (
	f "fmt"
)

func imprimirAprovados(aprovados ...string) {
	f.Println("Lista de Aprovados:")

	for i, aprovados := range aprovados {
		f.Printf("%d) %s \n", i+1, aprovados)
	}
}

func main() {
	// slice de aprovados
	aprovados := []string{"Maria", "Pedro", "Rafael", "Marta"}

	// os três pontos utilizados após a variável aprovados explode
	// o conteúdo do slice como se fossem parâmetros para a função de
	// imprimirAprovados.
	imprimirAprovados(aprovados...)

	// o código abaixo dá erro de compilação:
	// imprimirAprovados(aprovados)
	// Também não funciona com arrays:
	// array_aprovados := [...]string{"Maria", "Pedro", "Rafael", "Marta"}
	// imprimirAprovados(array_aprovados...)
}
