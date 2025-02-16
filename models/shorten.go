package models

import (
	"database/sql"
	"errors"
	"log"
)

type URLStore struct {
	LongURL  string
	ShortURL string
}

type ShortenerDataModel struct {
	DB *sql.DB
}

func (m *ShortenerDataModel) GetByShortURL(shortURL string) (*URLStore, error) {
	stmt := `SELECT original_url, shortened_url FROM urls
	WHERE shortened_url = ?`

	url := &URLStore{}
	err := m.DB.QueryRow(stmt, shortURL).Scan(&url.LongURL, &url.ShortURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}
	return url, nil
}

func (m *ShortenerDataModel) Insert(original string, shortened string) error {
	stmt := `INSERT OR REPLACE INTO urls (original_url, shortened_url) VALUES(?, ?)`
	_, err := m.DB.Exec(stmt, original, shortened)
	if err != nil {
		log.Printf("Error inserting into DB: %v", err)
	}
	return nil
}
