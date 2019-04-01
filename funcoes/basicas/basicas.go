package main

import (
	f "fmt"
)

func f1() {
	f.Println("Teste.")
}

func f2(p1 string, p2 string) {
	f.Printf("F2: %s %s \n", p1, p2)
}

// o tipo de retorno da função é uma string.
func f3() string {
	return "F3"
}

// como esses parâmetros são do mesmo tipo podemos fazer esse tipo de declaração
func f4(p1, p2 string) string {
	return f.Sprintf("F4: %s %s", p1, p2)
}

func f5(p1, p2 string, a1 float64) string {
	return f.Sprintf("F5: %s %s %.2f", p1, p2, a1)
}

// retorna duas strings | retorno múltiplos
func f6() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("e", "f")
	f.Println(f3())
	f.Println(f4("c", "d"))
	f.Println(f5("a", "b", 1.50))
	f.Println(f6())

	f.Println("-------------------")

	r3, r4 := f3(), f4("param1", "param2")
	f.Println(r3)
	f.Println(r4)

	f.Println("-------------------")

	// é possível ignorar qualquer retorno da função, para isso
	// utilizamos o underline, ex: r6, _ := f6()
	// neste exemplo acima ele está ignorando o segundo retorno.
	r61, r62 := f6()
	f.Println(r61)
	f.Println(r62)

	// ignora todos os retornos das funções
	f6()
}
