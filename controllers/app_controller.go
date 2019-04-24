package controllers

import (
	"net/http"

	"github.com/diegoajv/ChitChat/utils"
)

// Index HandleFunction
func Index(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, "", "layout", "navbar", "index")
}
