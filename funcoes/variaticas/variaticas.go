package main

import (
	f "fmt"
)

// funções variáticas são todas as funções que possuem uma quantidade
// de parâmetros variáveis - não fixa.

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	// o error NaN é o mesmo do js: Not a Number.
	// se tirar o tratamento abaixo ele retorna um NaN.
	if len(numeros) == 0 {
		return 0
	}

	return total / float64(len(numeros))
}

func main() {
	resultado := media(5.2, 5.5, 10, 8.9, 7.5)
	f.Printf("Média: %.2f \n", resultado)

	// inclusive é possível passar nenhum parâmetro.
	resultado = media()
	f.Printf("Média: %.2f \n", resultado)
}
