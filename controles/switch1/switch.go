package main

import (
	f "fmt"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		// o break já está incluso, ou seja, o switch é um if else
		fallthrough
		// fallthrough é o contrário de break, ou seja, diz pro Go
		// que ele deve seguir testando os outros cases.
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	f.Println(notaParaConceito(10))
	f.Println(notaParaConceito(9.8))
	f.Println(notaParaConceito(6.9))
	f.Println(notaParaConceito(2.1))
	f.Println(notaParaConceito(11))
	f.Println(notaParaConceito(-1.1))
}
