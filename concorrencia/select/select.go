package main

import (
	f "fmt"
	"time"

	"github.com/anitachiele/go_src/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.GetTitulo(url1)
	c2 := html.GetTitulo(url2)
	c3 := html.GetTitulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}

	/*
		Se deixar / descomentar o default ele sempre irá retornar o default pq
		é a única solução possível já as respostas ainda não estão presentes.

	*/
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	f.Println(campeao)
}
