package router

import (
	"net/http"

	"github.com/diegoajv/ChitChat/controllers"
)

// GetRouter returns app routes
func GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Configure static files
	files := http.FileServer(http.Dir("public"))
	router.Handle("/static/", http.StripPrefix("/static/", files))

	// App routes
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/login", controllers.LogIn)

	return router
}
