package main

import (
	f "fmt"
)

func main() {
	i := 1

	// criação de ponteiro (64bits)
	var p *int = nil

	// pega o endereço de memória da variável i
	p = &i
	f.Println("Endereço de memória p: ", p)
	f.Println("Endereço de memória i: ", &i)

	// *p == ao valor que o ponteiro aponta.
	f.Println("\np 01: ", *p)
	f.Println("i 01: ", i)
	*p++

	f.Println("\np 02: ", *p)
	f.Println("i 02: ", i)
	i++

	f.Println("\nDepois do i++:")

	f.Println("\np 03: ", *p)
	f.Println("i 03: ", i)

	// Go não tem aritmética de ponteiros: p++
}
