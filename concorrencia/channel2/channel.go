package main

import (
	f "fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	// os time.Sleep foram adicionado no código abaixo
	// apenas para não executarem muito rápido, para um
	// melhor entendimento.

	time.Sleep(time.Second)
	// enviando os dados para o channel
	c <- 2 * base

	time.Sleep(time.Second)
	// enviando os dados para o channel
	c <- 3 * base

	time.Sleep(3 * time.Second)
	// enviando os dados para o channel
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	f.Println("A")

	// recebendo os dados do channel
	a, b := <-c, <-c

	f.Println("B")
	f.Println(a, b)
	f.Println(<-c)
}
