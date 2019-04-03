package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			num := numeroAleatorio()

			time.Sleep(time.Second)
			c <- f.Sprintf("%s falando: %s", pessoa, getFrase(num))
		}
	}()

	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()

	return c
}

func getFrase(i int) string {
	frases := make(map[int]string)

	frases[0] = "Acho que vou pedir transferência."
	frases[1] = "Como vai você?"
	frases[2] = "Vou bem e você?"
	frases[3] = "Vai se mudar?"
	frases[4] = "Sim, vou pra outro planeta."
	frases[5] = "Tchau!"

	return frases[i]
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(5)
}

func main() {
	c := juntar(
		falar("Jãoa"),
		falar("Maria"),
	)

	f.Println(<-c, "|", <-c)
	f.Println(<-c, "|", <-c)
	f.Println(<-c, "|", <-c)
}
