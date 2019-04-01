package main

import (
	f "fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
// a sintaxy é:
// func (variavel nome_struct) nome_metodo() tipo_retorno
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto

	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}

	f.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	f.Println(produto2, produto2.precoComDesconto())
}
