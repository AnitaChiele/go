package main

// Não existem operadores ternários em GO

import (
	f "fmt"
)

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"

	// ternários:
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	f.Println(obterResultado(6.2))
}
