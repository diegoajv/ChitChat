package main

import (
	"net/http"

	"github.com/diegoajv/ChitChat/controllers"
)

func main() {
	// Configure app router
	router := http.NewServeMux()

	router.HandleFunc("/", controllers.Index)

	// Configure app server
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	server.ListenAndServe()
}
