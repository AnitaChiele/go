package main

import f "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	// ponteiro de int:
	var e *int

	// em go o NULL é nil.
	f.Printf("\n%v | %v | %v | %v | %v", a, b, c, d, e)

	// a string não é nula, é vazia.
	f.Printf("\n%v | %v | %v | %q | %v", a, b, c, d, e)
}
