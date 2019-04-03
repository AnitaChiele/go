package main

import (
	f "fmt"
)

/*
	Closure é uma função que envolve outra função - vulgo funções dentro de
	funções.

	Closure obviamente respeita seus escopos de criações das variáveis
	como será demonstrado neste arquivo:
*/

func closure() func() {
	x := 10

	var funcao = func() {
		f.Println(x)
	}

	return funcao
}

func main() {
	x := 20
	f.Println(x)

	imprimeX := closure()
	imprimeX()
}
