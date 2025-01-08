package main

import (
	"fmt"
	"net/http"

	"urlShortener/routes"
)

func main() {
	router := routes.InitRoutes()

	fmt.Println("Server started at:http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
