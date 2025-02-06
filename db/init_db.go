package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database by executing the schema.sql script.
func InitDB(dbPath, sqlFile string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	sqlStatements := string(sqlBytes)

	_, err = db.Exec(sqlStatements)
	if err != nil {
		log.Fatalf("Error executing SQL: %v", err)
	}

	fmt.Println("Database initialized successfully")
}
