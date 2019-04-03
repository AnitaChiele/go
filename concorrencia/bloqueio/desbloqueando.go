package main

import (
	f "fmt"
	"time"
)

func rotinaDesbloq(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	c <- 2 // operação bloqueante
	f.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer.

	go rotinaDesbloq(c)

	f.Println(<-c) // operação bloqueante.
	f.Println("Foi lido")
	f.Println(<-c)
	f.Println("Fim.")
}
