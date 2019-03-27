package main

import (
	f "fmt"
	"math"
	"reflect"
)

func main() {
	f.Println("\n***** Inteiros *****")
	// números inteiros
	f.Println(1, 2, 1000)

	// retorna o tipo da variável
	f.Println("Literal inteiro é:", reflect.TypeOf(3200))

	// sem sinal (só os positivos) uint8 uint16 uint32 uint64
	// byte é alias para o uint8
	var b byte = 255
	f.Println("O byte é:", reflect.TypeOf(b))

	// varia de arquitetura 64 ou 32 bits.
	i1 := math.MaxInt64
	f.Println("O valor máximo do int é:", i1)
	f.Println("O tipo de i1:", reflect.TypeOf(i1))

	// representa um mapeamento da tabela unicode (int32)
	var i2 rune = 'a'
	f.Println("O rune é: ", reflect.TypeOf(i2))
	// valor referente da tabela unide para a letra a
	f.Println("variável i2: ", i2)

	f.Println("\n\n***** Float *****")
	// números reais (float32 e float64)
	var x float32 = 49.99
	f.Println("O tipo de x é:", reflect.TypeOf(x))
	f.Println("O tipo default de 49.99 é:", reflect.TypeOf(49.99))

	var y = float32(49.99)
	f.Println("O tipo de y é:", reflect.TypeOf(y))

	f.Println("\n\n***** Boolean *****")
	// boolean
	bo := true
	f.Println("O tipo de bo é:", reflect.TypeOf(bo))
	f.Println("Valor negado de bo:", !bo)

	f.Println("\n\n***** String *****")
	// string
	s1 := "Olá meu nome é Leo"
	f.Println(s1 + "!")
	f.Println("O tamanho da string é:", len(s1))

	f.Println("\n\n***** String com múltiplas linhas *****")
	// string com múltiplas linhas
	// semelhante a múltiplas linhas em JS
	s2 := `Olá
	meu
	nome
	é
	Leo`
	f.Println("O tamanho da string s2 é:", len(s2))

	f.Println("\n\n***** Char *****")
	// char - go não possui tipo char.
	char := 'a'
	f.Println("O tipo de char é: ", reflect.TypeOf(char))
	// que é o código unicode da tabela - como já mostrado anteriormente.
	f.Println("O valor é do char: ", char)

	char2 := "a"
	f.Println("O tipo de char2 é: ", reflect.TypeOf(char2))
	f.Println("O valor é do char2: ", char2)
}
