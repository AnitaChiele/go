package main

import (
	"log"
	"net/http"
)

func main() {
	// o fileserver vai ler a partir do diretório public.
	fs := http.FileServer(http.Dir("public"))
	// quando tiver uma requisição de raiz da minha aplicação,
	// passe automaticamente para o fs.
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
