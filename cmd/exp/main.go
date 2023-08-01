package main

import (
	"database/sql"
	"fmt"

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

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      name TEXT,
      email TEXT UNIQUE NOT NULL
    );

    CREATE TABLE IF NOT EXISTS orders (
      id SERIAL PRIMARY KEY,
      user_id INT NOT NULL,
      amount INT,
      description TEXT
    );
    `)

	if err != nil {
		panic(err)
	}

	fmt.Println("Tables created.")

	userID := 1

	for i := 1; i <= 5; i++ {
		amount := i * 10
		description := fmt.Sprintf("fake order #%d", i)
		_, err := db.Exec(`
      INSERT INTO orders (user_id, amount, description)
      VALUES ($1, $2, $3)`, userID, amount, description)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Created fake orders")
}
