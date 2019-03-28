package main

import (
	f "fmt"
	"time"
)

// tipo interface permite trabalhar com tipo de dados
// de forma genérica.
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	f.Println(tipo(24.99))
	f.Println(tipo(12))
	f.Println(tipo("abc"))
	f.Println(tipo(func() {}))
	f.Println(tipo(time.Now()))
}
