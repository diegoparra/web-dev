package controllers

import (
	"html/template"
	"net/http"

	"github.com/diegoparra/calhoun/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	questios := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, it has a free version.",
		},
		{
			Question: "What are the support hours?",
			Answer:   "we have support staff answering email 27/7",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:diegoparraferreira@gmail.com">support@dpferreira.com.br</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questios)
	}
}
