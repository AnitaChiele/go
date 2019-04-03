package main

import (
	f "fmt"

	"github.com/anitachiele/go_src/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.GetTitulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.GetTitulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	f.Println(<-c, "|", <-c)
	f.Println(<-c, "|", <-c)
}
