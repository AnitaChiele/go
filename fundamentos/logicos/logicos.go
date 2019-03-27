package main

import (
	f "fmt"
)

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // não existe o OU exclusivo em GO
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	f.Printf("\nTv50: %t | Tv32: %t | Saudável: %t",
		tv50, tv32, !sorvete)

	tv50, tv32, sorvete = compras(true, false)

	f.Printf("\nTv50: %t | Tv32: %t | Saudável: %t",
		tv50, tv32, !sorvete)

	tv50, tv32, sorvete = compras(false, false)

	f.Printf("\nTv50: %t | Tv32: %t | Saudável: %t",
		tv50, tv32, !sorvete)
}
