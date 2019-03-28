package main

import (
	f "fmt"
	"reflect"
)

func main() {
	// array
	a1 := [3]int{1, 2, 3}

	// slice - tem tamanho variável.
	// não é um array, ele é um trecho de um array,
	// que utiliza ponteiro para apontar pro primeiro
	// elemento.
	s1 := []int{1, 2, 3}

	f.Println(a1, s1)
	f.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	f.Println(a2, s2)
	f.Println("-----------------")

	a3 := [5]int{6, 7, 8, 9, 10}

	s3 := a3[1:3]
	f.Println(s3)

	f.Println("-----------------")

	s4 := a3[:2]
	f.Println(s4)

	s4 = a3[:3]
	f.Println(s4)

	s5 := s4[:1]
	f.Println(s5)
}
