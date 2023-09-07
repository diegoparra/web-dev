package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/diegoparra/calhoun/controllers"
	"github.com/diegoparra/calhoun/models"
	"github.com/diegoparra/calhoun/templates"
	"github.com/diegoparra/calhoun/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		fmt.Println("error from db.Open")
		panic(err)
	}

	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))

	usersC.Templates.SignIn = views.Must(
		views.ParseFS(templates.FS, "signin.html", "tailwind.html"),
	)
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/sigin", usersC.SignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server on Port: 3000")
	http.ListenAndServe(":3000", r)
}
