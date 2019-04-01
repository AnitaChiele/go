package main

import (
	f "fmt"
)

/*
	A função init inicializa mesmo sem ter sido chamada e antes do método main.
*/

func init() {
	f.Println("Inicializando...")
}

func main() {
	f.Println("Main")
}
