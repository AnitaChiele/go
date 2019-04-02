package main

import (
	"encoding/json"
	f "fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899.99,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	p1Json, _ := json.Marshal(p1)
	f.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":2.40,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	f.Println(p2.Tags[1])
}
