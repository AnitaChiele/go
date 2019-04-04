package main

import (
	"database/sql"
	"fmt"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hashPass := fmt.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	linhas, _ := res.RowsAffected()

	fmt.Println("id: ", id, "linhas alteradas: ", linhas)
}
