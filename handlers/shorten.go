package handlers

import (
	"fmt"
	"net/http"

	"urlShortener/short"
)

var urlStore = make(map[string]string)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")

	shortURL := short.ShortCode()

	urlStore[shortURL] = longURL

	fmt.Fprintf(w, "http://localhost:8080/%s\n", shortURL)
}
