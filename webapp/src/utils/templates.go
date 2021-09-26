package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//LoadTemplates inserts the templates in the var templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//ExecTemplate render a html in the page
func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
