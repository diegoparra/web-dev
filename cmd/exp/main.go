package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/diegoparra/calhoun/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		fmt.Println("error from db.Open")
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("diego2@parra.com", "123")
	if err != nil {
		fmt.Println("Error from create")
		panic(err)
	}

	fmt.Println(user)
}
