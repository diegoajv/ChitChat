package controllers

import (
	"html/template"
	"net/http"
)

// LogIn HandleFunction
func LogIn(writer http.ResponseWriter, request *http.Request) {
	loginTemplate := "templates/login.html"

	templates := template.Must(template.ParseFiles(loginTemplate))
	templates.Execute(writer, "")
}
