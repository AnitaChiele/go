package main

import (
	f "fmt"

	"github.com/anitachiele/go_src/titulo"
)

/*
	Google I/O 2012 - Go Concurrency Patterns:
	https://www.youtube.com/watch?v=f6kdp27TYZs
*/

func main() {
	t1 := titulo.GetTitulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo.GetTitulo("https://www.amazon.com", "https://www.youtube.com")

	f.Println("Primeiros: ", <-t1, "|", <-t2)
	f.Println("Segundos: ", <-t1, "|", <-t2)
}
