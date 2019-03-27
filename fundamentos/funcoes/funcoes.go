package main

/*
	Para executar esse código é necessário rodá-lo através do terminal.
	Entrar nesta pasta go/funcoes

	No linux:
	go run *.go

	No windows:
	go run funcoes.go main.go
*/

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
