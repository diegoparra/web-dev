package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Diego",
		Age:  32,
	}

	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, user)
}
