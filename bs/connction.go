package bs

import (
	"database/sql"
	"fmt"
)

var dbds *sql.DB

func Connection() {

	var err error
	dbds, err = sql.Open("mysql", "root:Srinath@1608@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	RegisterAccountHandlers()
	defer dbds.Close()

}
