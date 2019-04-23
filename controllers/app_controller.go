package controllers

import (
	"html/template"
	"net/http"
)

// Index HandleFunction
func Index(writer http.ResponseWriter, request *http.Request) {
	indexTemplate := "templates/index.html"

	templates := template.Must(template.ParseFiles(indexTemplate))
	templates.Execute(writer, "")
}
