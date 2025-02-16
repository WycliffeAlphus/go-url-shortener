package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database by executing the schema.sql script.
func InitDB(dbPath, sqlFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {

		return nil, errors.New("Error opening database: " + err.Error())
	}

	defer db.Close()

	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		return nil, errors.New("Error reading SQL file: " + err.Error())
	}
	sqlStatements := string(sqlBytes)

	_, err = db.Exec(sqlStatements)
	if err != nil {
		return nil, errors.New("Error executing SQL: " + err.Error())
	}

	fmt.Println("Database initialized successfully")
	return db, nil
}
