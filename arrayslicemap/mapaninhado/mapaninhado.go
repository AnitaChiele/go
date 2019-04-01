package main

import (
	f "fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float32{
		"G": {
			"Gabriela Silva": 15456.1,
			"Guga Pereira":   5456.1,
		},
		"J": {
			"José João": 2000.00,
		},
		"P": {
			"Pedro Joao": 1200.00,
		},
	}

	f.Println(funcsPorLetra)
	f.Println("\n")

	for letra, funcs := range funcsPorLetra {
		f.Println(letra, funcs)
	}

	delete(funcsPorLetra, "P")

	f.Println("\n")

	for letra, funcs := range funcsPorLetra {
		f.Println(letra, funcs)
	}
}
