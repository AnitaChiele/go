package main

import (
	f "fmt"
)

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	}

	return "E"
}

func main() {
	f.Println(notaParaConceito(9.8))
	f.Println(notaParaConceito(6.9))
	f.Println(notaParaConceito(2.1))
}
