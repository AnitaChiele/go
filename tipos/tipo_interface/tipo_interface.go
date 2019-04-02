package main

import (
	f "fmt"
)

type curso struct {
	nome string
}

func main() {
	/*
		Existem duas formas de criar variáveis de tipagem dinâmicas:

		Criando um tipo de interface vazia e declarando a variável com esse tipo
		criado. (coisa)

		Declarando a variável como do tipo interface vazia. (coisa2)
	*/
	type dinamico interface{}
	var coisa dinamico = "Opa"
	f.Println("coisa:", coisa)

	coisa = 1
	f.Println("coisa:", coisa)

	coisa = true
	f.Println("coisa:", coisa)

	coisa = curso{"Golang: Explorando a linguagem do google"}
	f.Println("coisa:", coisa)

	var coisa2 interface{}
	coisa2 = "Opa"
	f.Println("coisa 02:", coisa2)

	coisa2 = 1
	f.Println("coisa 02:", coisa2)

	coisa2 = true
	f.Println("coisa 02:", coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do google"}
	f.Println("coisa 02:", coisa2)

}
