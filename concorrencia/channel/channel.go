package main

import (
	f "fmt"
)

// o channel é um tipo de dado assim como o int, float e etc.
func main() {
	// criando um channel que vai passar tipos ints.
	// 1 é o buffer.
	ch := make(chan int, 1)

	ch <- 1 // está enviando o valor 1 para o channel (escrita).
	<-ch    // recebendo dados do channel (leitura).

	ch <- 2
	f.Println(<-ch)

	/*
		O channel é um mecanismo de sincronismo para trabalhar com
		as Go Routines, ou seja, o channel é um ponto de sincronismo em
		GO.

		O channel possibilita que a execução fique esperando o retorno
		de algum valor para depois continuar a execução padrão do código.
		Lembrando muito a execução assíncrona dos programas.
	*/
}
