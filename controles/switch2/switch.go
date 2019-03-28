package main

import (
	f "fmt"
	"time"
)

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	t := time.Now()

	// switch true
	switch {
	case t.Hour() < 12:
		f.Println("Bom dia!")
	case t.Hour() < 18:
		f.Println("Boa tarde.")
	default:
		f.Println("Boa noite.")
	}

	f.Println(notaParaConceito(9.8))
	f.Println(notaParaConceito(6.9))
	f.Println(notaParaConceito(2.1))

}
