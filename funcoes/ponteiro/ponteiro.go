package main

import (
	f "fmt"
)

func inc1(n int) {
	n++
}

// um ponteiro é representado por *
func inc2(n *int) {
	// desreferencia um ponteiro: *[ponteiro]
	*n++
}

func main() {
	n := 1

	inc1(n)
	f.Println(n)

	inc2(&n) // passando por referência
	f.Println(n)
}
