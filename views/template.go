package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTmlp *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		return Template{}, fmt.Errorf("Pasing template: %w", err)
	}

	return Template{
		htmlTmlp: tpl,
	}, err
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTmlp.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
	}
}
