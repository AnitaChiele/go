package main

import (
	"database/sql"
	"fmt"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	// o final da senha deve ser marcado pelo @/
	hashPass := fmt.Sprintf("root:%s@/", local.GetSenha())
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `CREATE TABLE usuarios (
			id INTEGER AUTO_INCREMENT,
			nome VARCHAR(100),
			PRIMARY KEY (id)
	)`)
}
