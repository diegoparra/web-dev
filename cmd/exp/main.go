package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open(
		"pgx",
		"host=localhost port=5432 user=root password=root dbname=lenslocked sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
