package main

import (
	f "fmt"
)

func main() {
	// já declarado e inicializado o map.
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	f.Println(funcsESalarios)

	funcsESalarios["Rafael Filho"] = 1350.0

	f.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistente")

	f.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		f.Println(nome, salario)
	}
}
