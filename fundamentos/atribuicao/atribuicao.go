package main

import f "fmt"

func main() {
	var b byte = 3
	f.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	f.Println(i)

	// atribuição de múltiplos valores
	x, y := 1, 2
	f.Println(x, y)

	// dispensa a criação de variável auxiliar quando
	// se realiza a troca de valor entre elas.
	x, y = y, x
	f.Println(x, y)
}
