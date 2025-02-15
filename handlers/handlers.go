package handlers

import (
	"net/http"
	"text/template"

	"urlShortener/models"
)

type PageData struct {
	BaseURL string
	Error   string
	URLData []*models.URLStore
}

type App struct {
	urls *models.ShortenerDataModel
}

func NewApp(urls *models.ShortenerDataModel) *App {
	return &App{
		urls: urls,
	}
}

func (app *App) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		urls, err := app.urls.Latest()

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			BaseURL: "http://" + r.Host,
			URLData: urls,
		}
		tmpl, err := template.ParseFiles("templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	}
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	longURL, ok := urlStore[shortURL]

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
