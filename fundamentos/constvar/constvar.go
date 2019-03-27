package main

import (
	"fmt"
	m "math"
)

// o m na frente do math é um alias que pode ser utilizado
// no lugar de math.
// também é possível utilizar _ caso tenha que importar algo
// que não seja utilizado diretamente no arquivo:
// import _ "nomeDoImport"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador
	// é possível criar a variável utilizado a sigla var no início,
	// ou podemos utilizar a forma reduzida := (utilizada na variável
	// area abaixo).

	// forma reduzida de criar uma variável
	// todas as variáveis tem que ser utilizadas em GO, caso contrário
	// dá um erro de compilação.
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	// ao declarar e atribuir a variável é necessário utilizar :=
	// só será possível utilizar = quando for atualizar o valor da
	// variável. Lembrando que o tipo da variável não pode ser alterado.

	// declaração de constantes
	const (
		a = 1
		b = 2
	)

	// declaração de variáveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	// sintaxy reduzida:
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)

	teste := 1
	fmt.Println(teste)
	// não é possível trocar o tipo da variável:
	// teste = 2.5

	// os := são utilizados para criar uma variável, ou seja, caso queria
	// alterar o valor dela o exemplo abaixo não funciona (erro de compilação):
	// teste := 2

	teste = 2
	fmt.Println(teste)
}
