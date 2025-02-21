package main

import (
	"log"
	"net/http"

	"urlShortener/db"
	"urlShortener/handlers"
	"urlShortener/middlewares"
	"urlShortener/models"
	"urlShortener/routes"
)

func main() {
	dbPath := "/mnt/volume/db/database.sqlite3"
	sqlFile := "db/schema.sql"

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

	log.Println("Server started at: http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
