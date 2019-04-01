package main

import (
	f "fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, f.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	_, err := fatorial(-5)

	if err != nil {
		f.Println(err)
	}

	f.Println("Resultado: ", resultado)

}
