package main

import (
	f "fmt"
)

func main() {
	// map[tipo da chave] tipo do valor:
	// var aprovados map[int]string
	// mapas sempre devem ser inicializados, caso contrário, dará erro de compilação.
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[98765432] = "Pedro"
	aprovados[65439102] = "Você"

	f.Println(aprovados)

	for cpf, nome := range aprovados {
		f.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	f.Println(aprovados[65439102])

	delete(aprovados, 98765432)
	f.Println(aprovados[98765432])
	f.Println(aprovados)
}
