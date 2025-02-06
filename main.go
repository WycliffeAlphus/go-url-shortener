package main

import (
	"fmt"
	"net/http"

	"urlShortener/db"
	"urlShortener/routes"
)

func main() {
	router := routes.InitRoutes()
	dbPath := "db/database.sqlite3"
	sqlFile := "db/schema.sql"

	db.InitDB(dbPath, sqlFile)

	fmt.Println("Server started at:http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
