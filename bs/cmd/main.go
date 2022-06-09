package main

import (
	"atms/bs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("stating atms")
	bs.Connection()

}
