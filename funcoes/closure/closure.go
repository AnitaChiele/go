package main

import (
	f "fmt"
)

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
