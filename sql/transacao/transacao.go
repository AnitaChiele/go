package main

import (
	"database/sql"
	f "fmt"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hashPass := f.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(?,?)")
	stmt.Exec(4, "Bia")
	stmt.Exec(5, "Carlos")
	// a chave duplicada abaixo é de propósito para elucidar uma transação onde 1 falha
	// e todas as outras execuções de query são feitas Rollback.
	// ou seja, nenhuma das transações é aplicada.
	_, err = stmt.Exec(1, "Thiago")

	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// }

	tx.Commit()
}
