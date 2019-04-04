package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	hashPass := fmt.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", 1)
	// fecha o resultset
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
