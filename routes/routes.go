package routes

import (
	"net/http"

	"urlShortener/handlers"
)

// InitRoutes initializes the routes for the application
func InitRoutes(app *handlers.App) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HandleHome())
	mux.HandleFunc("/shorten", app.HandleShorten())
	mux.HandleFunc("/redirect", app.HandleRedirect)
	return mux
}
