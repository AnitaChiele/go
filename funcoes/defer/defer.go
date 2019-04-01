package main

import (
	f "fmt"
)

// defer faz com que a instrução seja executada
// no momento mais tardio possível dentro do seu
// escopo.

func obterValorAprovado(numero int) int {
	defer f.Println("fim!")

	if numero >= 5000 {
		f.Println("Valor alto...")
		return 5000
	}

	f.Println("Valor baixo..")
	return numero
}

func main() {
	f.Println(obterValorAprovado(6000))
	f.Println(obterValorAprovado(3000))
}
