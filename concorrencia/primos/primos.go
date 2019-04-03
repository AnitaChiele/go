package main

import (
	f "fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2

	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1

				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	// fechando o channel
	close(c)
}

func main() {
	c := make(chan int, 30)
	// cap(c) retorna a capacidade do channel - buffer.

	go primos(60, c)

	for primo := range c {
		f.Printf("%d ", primo)
	}

	f.Println("Fim!")
}
