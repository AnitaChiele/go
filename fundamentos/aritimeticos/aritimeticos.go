package main

import (
	f "fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	f.Println("Soma =>", a+b)
	f.Println("Subtração =>", a-b)
	f.Println("Divisão =>", a/b)
	f.Println("Multiplicação =>", a*b)
	f.Println("Módulo =>", a%b)

	// operações binárias
	f.Println("\n\n----------------------------")
	f.Println("Operações Binárias:")
	f.Println("E =>", a&b)
	f.Println("Ou =>", a|b)
	f.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	// Operações utilizando a biblio math
	f.Println("\n\n----------------------------")
	f.Println("Operações utilizando Math:")
	f.Println("Maior =>", math.Max(float64(a), float64(b)))
	f.Println("Menor =>", math.Min(c, d))
	f.Println("Exponenciação =>", math.Pow(c, d))
}
