package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Teste12331", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
