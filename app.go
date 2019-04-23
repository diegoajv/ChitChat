package main

import (
	"net/http"

	"github.com/diegoajv/ChitChat/router"
)

func main() {
	// Configure app server
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router.GetRouter(),
	}

	server.ListenAndServe()
}
