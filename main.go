package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	"github.com/diegoparra/calhoun/controllers"
	"github.com/diegoparra/calhoun/migrations"
	"github.com/diegoparra/calhoun/models"
	"github.com/diegoparra/calhoun/templates"
	"github.com/diegoparra/calhoun/views"
)

func main() {
	// Setup database
	cfg := models.DefaultPostgresConfig()
	fmt.Println(cfg.String())
	db, err := models.Open(cfg)
	if err != nil {
		fmt.Println("error from db.Open")
		panic(err)
	}

	defer db.Close()

	err = models.MigrateFs(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Start services
	userService := models.UserService{
		DB: db,
	}

	sessionService := models.SessionService{
		DB: db,
	}

	// Setup middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := "gg2eCDG7Uioy9OJKfiuuWLfwOUNE2ZQq"
	csrfMiddle := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	// Setup controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))

	usersC.Templates.SignIn = views.Must(
		views.ParseFS(templates.FS, "signin.html", "tailwind.html"),
	)

	// Setup router and routes

	r := chi.NewRouter()
	r.Use(csrfMiddle)
	r.Use(umw.SetUser)
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	// r.Get("/users/me", usersC.CurrentUser)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// Start server
	fmt.Println("Starting server on Port: 3000")
	http.ListenAndServe(":3000", r)
}
