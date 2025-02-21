package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"urlShortener/models"
	"urlShortener/short"
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
		data := PageData{
			BaseURL: "http://" + r.Host,
		}
		tmpl, err := template.ParseFiles("templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	}
}

func (app *App) HandleShorten() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Ensure response is always JSON

		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var input struct {
			URL string `json:"url"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
			return
		}

		if input.URL == "" {
			http.Error(w, `{"error":"URL is required"}`, http.StatusBadRequest)
			return
		}

		shortCode := short.ShortCode()
		err := app.urls.Insert(input.URL, shortCode)
		if err != nil {
			http.Error(w, `{"error":"Error creating short url"}`, http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"original_url": input.URL,
			"short_url":    "http://" + r.Host + "/" + shortCode,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (app *App) HandleRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		if shortURL == "" {
			http.NotFound(w, r)
			return
		}
		url, err := app.urls.GetByShortURL(shortURL)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url.LongURL, http.StatusFound)
	}
}
