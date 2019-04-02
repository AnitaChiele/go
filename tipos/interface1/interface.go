package main

import (
	f "fmt"
)

/*
	Em GO ao implementar interfaces não é preciso dizer explicitamente que as estruturas
	implementam a interface. Para que isso aconteça, as estruturas tem que implementar os
	métodos declarados na interface. Caso seja mais de um método, então ela deverá aplicar
	todos os métodos.

	Em tempo de compilação o compilador valida se as estruturas implementam todos os métodos
	da interface e relaciona eles.
*/

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return f.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

// recebe por parâmetro qualquer coisa que respeita a interface
// imprimivel
func imprimir(x imprimivel) {
	f.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	f.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{
		nome:  "Calça Jeans",
		preco: 79.90,
	}
	imprimir(coisa)

	imprimir(produto{
		nome:  "p2",
		preco: 39.90,
	})
}
