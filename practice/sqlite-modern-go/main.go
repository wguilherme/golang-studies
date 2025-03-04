package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./test.db")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

}
