package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// a quebra de linha é no final da string.
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// não dá pra utilizar a concatenação no print para
	// imprimir uma variável que não seja do tipo string,
	// ex: fmt.Println("O valor de x é "+ x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x, "!!!")

	// print formatado
	fmt.Printf("O valor de x é %f \n", x)
	// pega só as últimas duas casas decimais:
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// %v é genérico e ele imprime todo o tipo de varíavel
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
