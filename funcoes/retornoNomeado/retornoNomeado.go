package main

import (
	f "fmt"
)

/*
	Essa função utilizará o retorno limpo. Ou seja, não é necessário
	informar a ordem do retorno já que ela está declarada na função
	- após a declaração dos parâmetros.
	Caso queira trocar a ordem do retorno, basta declarar o return
	normalmente. (return p1, p2)
*/
func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return
}

/*
	Não é possível executar o código abaixo porque ele dá erro de compilação
	alegando que as variáveis estão duplicadas:

	func trocar(p1, p2 int) (p2 int, p1 int) {
		return
	}
*/

func main() {
	primeiro := 1
	segundo := 2

	f.Println(primeiro, segundo)
	f.Println(trocar(primeiro, segundo))
}
