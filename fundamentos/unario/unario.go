package main

import (
	f "fmt"
)

func main() {
	x := 1
	y := 2

	// apenas posfix
	x++ // x +=1 ou x = x + 1
	f.Println(x)

	y-- // y -=1 ou y = y + 1
	f.Println(y)
}
