package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anitachiele/config_db/local"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hashPass := fmt.Sprintf("root:%s@/%s", local.GetSenha(), "cursogo")
	db, err := sql.Open("mysql", hashPass)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Uóxiton Istive", 2)
	stmt.Exec("Sheristone Ualeska", 3)

	stmt, _ = db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt.Exec(5)

}
