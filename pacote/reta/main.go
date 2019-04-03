package main

import (
	f "fmt"
)

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	f.Println(catetos(p1, p2))
	f.Println(Distancia(p1, p2))
}
