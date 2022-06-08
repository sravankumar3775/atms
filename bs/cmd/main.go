package main

import (
	"atms/bs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Srinath@1608@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	bs.Connection(db)

	defer db.Close()
}
