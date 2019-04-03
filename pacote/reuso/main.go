package main

import (
	f "fmt"

	"github.com/anitachiele/areaantigo"
)

func main() {
	f.Println(areaantigo.Circ(6.0))
	f.Println(areaantigo.Rect(5.0, 2.0))

	// a função abaixo não é acessível pq está com visibilidade privada
	// para acessos fora do pacote em que foi criada:
	// f.Println(areaantigo._Triangulo(5.0, 2.0))
}
