package main

import (
	"fmt"
	"net/http"

	"urlShortener/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to URL shortener!!")
	})

	http.HandleFunc("/shorten", handlers.HandleShorten)
	http.HandleFunc("/redirect", handlers.HandleRedirect)

	fmt.Println("Server started at:http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
