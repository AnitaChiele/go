package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// o i só vai existir no if e else.
	// não será acessível ou visível fora dessa estrutura.
	// isso também funciona com switch
	if i := numeroAleatorio(); i < 5 {
		f.Println("Ganhou!!!", i)
	} else {
		f.Println("Perdeu!", i)
	}
}
