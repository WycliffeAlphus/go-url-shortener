package routes

import (
	"net/http"

	"urlShortener/handlers"
)

// InitRoutes initializes the routes for the application
func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handlers.HandleShorten)
	mux.HandleFunc("/redirect", handlers.HandleRedirect)
	return mux
}
