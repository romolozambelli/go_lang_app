package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Insert all the html template into variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Render the html pages
func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
