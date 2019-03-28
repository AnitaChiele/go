package main

import (
	f "fmt"
)

func main() {
	i := 1

	//  não tem while em Go
	for i <= 10 {
		f.Println(i)
		i++
	}

	// for está sendo incrementado de 2 em 2.
	for i := 0; i <= 20; i += 2 {
		f.Printf("%d ", i)
	}

	f.Println("\nMisturando..")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			f.Print("Par ")
		} else {
			f.Print("Impar ")
		}
	}

	// loop infinito
	// for {
	// 	f.Println("Para sempre...")
	// 	time.Sleep(time.Second * 5)
	// }

	// Rodar estruturalmente foreach será visto no capítulo de array.
}
