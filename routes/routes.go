package routes

import (
	"net/http"

	"urlShortener/handlers"
)

// InitRoutes initializes the routes for the application
func InitRoutes(app *handlers.App) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", app.HandleShorten())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			app.HandleHome()(w, r)
			return
		}
		app.HandleRedirect()(w, r) // redirects unknown paths
	})
	return mux
}
