package bs

import "database/sql"

var dbds *sql.DB

func Connection(acc *sql.DB) {
	dbds = acc
	RegisterAccountHandlers()
	//RegisterLoginHandlers()
}
