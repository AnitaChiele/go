package main

import (
	"database/sql"
	"encoding/json"
	f "fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

// Usuario :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para a função adequeada.
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		f.Fprintf(w, "Desculpa.. :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	hashPass := f.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	f.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	hashPass := f.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	f.Fprint(w, string(json))
}
