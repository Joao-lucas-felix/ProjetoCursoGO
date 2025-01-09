package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates load all templates of the pages
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}