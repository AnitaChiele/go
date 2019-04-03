package main

import (
	f "fmt"

	"github.com/anitachiele/go_src/html"
)

/*
	Google I/O 2012 - Go Concurrency Patterns:
	https://www.youtube.com/watch?v=f6kdp27TYZs

	Generator é um padrão. Para entender mais veja o package.
*/

func main() {
	t1 := html.GetTitulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := html.GetTitulo("https://www.amazon.com", "https://www.youtube.com")

	f.Println("Primeiros: ", <-t1, "|", <-t2)
	f.Println("Segundos: ", <-t1, "|", <-t2)
}
