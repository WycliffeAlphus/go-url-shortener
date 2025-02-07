package models

import (
	"database/sql"
)

type URLStore struct {
	LongURL, ShortURL string
	Clicks            int
}

type ShortenerDataModel struct {
	DB *sql.DB
}

// Latest returns the latest URLs
func (m *ShortenerDataModel) Latest() ([]*URLStore, error) {
	stmt := `SELECT original_url, shortened_url, clicks FROM urls1`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := []*URLStore{}

	for rows.Next() {
		url := &URLStore{}

		if err := rows.Scan(&url.LongURL, &url.ShortURL, &url.Clicks); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (m *ShortenerDataModel) Insert(original string, shortened string, clicks int) (int, error) {
	stmt := `INSERT INTO ulrs (original_url, shortened_url, clicks) VALUES(?, ?, ?)`
	result, err := m.DB.Exec(stmt, original, shortened, clicks)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}
