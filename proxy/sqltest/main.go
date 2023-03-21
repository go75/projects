package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	d, err := sql.Open("sqlite3", "./go.db")
	if err!=nil {
		panic(err)
	}
	defer d.Close()
	err = d.Ping()
	if err!=nil {
		panic(err)
	}
}