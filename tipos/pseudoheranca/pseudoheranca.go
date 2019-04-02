package main

import (
	f "fmt"
)

type carro struct {
	nome          string
	velodadeAtual int
}

/*
	O carro na estrutura ferrari é um campo anônimo.
	Ele possibilita "herdar" todos os atributos do tipo carro.
	Na verdade o nome disso é composição.
*/
type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	ferrari := ferrari{}
	ferrari.nome = "F40"
	ferrari.velodadeAtual = 0
	ferrari.turboLigado = true

	f.Printf("A ferrari %s está com o turbo ligado? %v\n", ferrari.nome, ferrari.turboLigado)
	f.Println(ferrari)
}
