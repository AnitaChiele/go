package main

import (
	f "fmt"
)

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

/*
	Para fazer alterações nos atributos das estruturas é necessário trabalhar com ponteiros.
	Essa função é só um exemplo, não é recomendado a utilização de interfaces para alterar
	conteúdo de métodos.

	Em geral as interfaces devem ser utilizadas apenas para lerem esses
	conteúdos de forma que não cause nenhum impacto nas estruturas.
*/
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// como o carro2 está sendo trabalhado a partir do nível da interface
	// é necessário passar o endereço de memória da estrutura da ferrari.
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	f.Println(carro1, carro2)
}
