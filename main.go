package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"urlShortener/db"
	"urlShortener/handlers"
	"urlShortener/middlewares"
	"urlShortener/models"
	"urlShortener/routes"
)

func main() {
	// Set up database paths for both development and production
	dbDir := "/app/database"
	if os.Getenv("GO_ENV") != "production" {
		dbDir = "db"
	}

	// Ensure database directory exists
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	dbPath := filepath.Join(dbDir, "database.sqlite3")
	sqlFile := filepath.Join("db", "schema.sql")

	dBase, err := db.InitDB(dbPath, sqlFile)
	if err != nil {
		log.Fatal(err)
	}

	urlModel := &models.ShortenerDataModel{DB: dBase}
	app := handlers.NewApp(urlModel)
	router := routes.InitRoutes(app)

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := middlewares.Logger(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at port :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
