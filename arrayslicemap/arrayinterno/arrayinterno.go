package main

import (
	f "fmt"
)

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	f.Println(s1, s2)

	s1[0] = 7
	f.Println(s1, s2)

	// ao atualizar o conteúdo da posição 0 do slice 01 para
	// 7 foi possível verificar com o print que ele também
	// aplicou para o s2, isso não significa que ele fez a alteração
	// nos slices de forma separada, mas indica e comprova que
	// os dois slices apontam para o mesmo array interno.
}
