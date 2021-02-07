package main

import (
	"database/sql"
	"fmt"
)

func main() {
	url := "root:@tcp(localhost:3306)/watch_test?parseTime=true&charset=utf8"
	db, err := sql.Open("mysql", url)
	fmt.Println(err)
	fmt.Println(db)
}
