package main

import (
	f "fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	f.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer.

	go rotina(c)

	f.Println(<-c) // operação bloqueante.
	f.Println("Foi lido")
	f.Println(<-c) // deadlock.
	f.Println("Fim.")
}
