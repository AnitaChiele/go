package main

import (
	f "fmt"
)

func main() {
	// como ao criar esse slice ele não foi baseado em nenhum array
	// ele criou um array e esse slice faz referência para um interno
	// tem um ponteiro para o elemento array e tem um tamanho de 10 elementos.
	s := make([]int, 10)
	s[9] = 12

	f.Println(s)

	// o 20 do make é sinalizando o array que ele vai criar internamente deverá ter
	// 20 elementos [capacidade máxima do array interno]
	s = make([]int, 10, 20)

	//cap é a capacidade máxima do array interno
	f.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	f.Println(s, len(s), cap(s))

	s = append(s, 1)
	f.Println(s, len(s), cap(s))

	// o Go gerencia a quantidade de elementos internos do array para nós,
	// ou seja, ao "estourar" a capacidade do array interno automaticamente
	// a própria linguagem gerencia isso, pega o array atual e copia para um com
	// tamanho dobrado e isso te dá mais flexibilidade para trabalhar com arrays.
}
