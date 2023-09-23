package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tmpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, r, nil)
	}
}

func FAQ(tmpl Template) http.HandlerFunc {
	questions := []struct {
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
		tmpl.Execute(w, r, questions)
	}
}
