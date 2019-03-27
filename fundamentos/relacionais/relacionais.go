package main

import (
	f "fmt"
	"time"
)

func main() {
	f.Println("==", "Banana" == "Banana")
	f.Println("!=", 3 != 2)
	f.Println("<", 3 < 2)
	f.Println(">", 3 > 2)
	f.Println("<=", 3 <= 2)
	f.Println(">=", 3 >= 2)

	f.Println("\n-------------------------------------")
	// objeto do tipo Data
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	// não compara os ponteiros
	f.Println("Datas:", d1 == d2)
	f.Println("Datas:", d1.Equal(d2))

	// structs
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	f.Println("Pessoas:", p1 == p2)

	p3 := Pessoa{"João"}
	p4 := Pessoa{"Joã"}
	f.Println("Pessoas:", p3 == p4)
}
