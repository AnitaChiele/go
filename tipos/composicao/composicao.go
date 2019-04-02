package main

import (
	f "fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// Se for necessário, é possível adicionar mais métodos sem perder
// a composição das interfaces.
type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	f.Println("Turbo!")
}

func (b bmw7) fazerBaliza() {
	f.Println("Estacionando...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}

	b.ligarTurbo()
	b.fazerBaliza()
}
