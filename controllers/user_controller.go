package controllers

import (
	"net/http"

	"github.com/diegoajv/ChitChat/utils"
)

// LogIn HandleFunction
func LogIn(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, "", "layout", "navbar", "login")
}
