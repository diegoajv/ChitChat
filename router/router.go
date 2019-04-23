package router

import (
	"net/http"

	"github.com/diegoajv/ChitChat/controllers"
)

// GetRouter creates a multiplexer
func GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/login", controllers.LogIn)

	return router
}
