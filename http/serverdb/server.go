package main

import (
	"log"
	"net/http"
)

/*
	Para executar:
	`$ go run *.go`
	Ou
	`$ go run cliente.go server.go`
*/
func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
