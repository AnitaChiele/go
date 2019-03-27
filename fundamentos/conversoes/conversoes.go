package main

import (
	f "fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// não é possível realizar operações entre tipos diferentes, ex:
	// f.Println(x / y) não compila, ou converte-se o inteiro pra float
	// ou o float pra inteiro:
	f.Println(x / float64(y))
	f.Println(int(x) / y)

	f.Println("\n\n-----------------------------------")

	nota := 6.9
	notaFinal := int(nota)
	f.Println(notaFinal)

	// o 123 é referente ao código da tabela ascii:
	f.Println("Teste: " + string(97))
	f.Println("Teste: ", string(123))

	f.Println("\nConvertendo int ==> String: ")
	// int para string:
	f.Println("Teste: " + strconv.Itoa(123))

	f.Println("\nConvertendo String ==> int: ")
	// String para int:
	num, _ := strconv.Atoi("123")

	// a função acima retorna o número e o erro.
	f.Println(num - 123)

	f.Println("\nConvertendo String ==> Boolean: ")
	// string para bool:
	b, _ := strconv.ParseBool("true")
	if b {
		f.Println("Vdd...")
	}

	b2, _ := strconv.ParseBool("1")
	if b2 {
		f.Println("Vdd...")
	}

	b3, _ := strconv.ParseBool("0")
	if b3 {
		f.Println("Vdd...")
	} else {
		f.Println("False..")
	}

	b4, _ := strconv.ParseBool("5")
	if b4 {
		f.Println("Vdd...")
	} else {
		f.Println("False..")
	}
}
