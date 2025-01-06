package handlers

import "net/http"

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	longURL, ok := urlStore[shortURL]

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
