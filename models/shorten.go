package models

import (
	"database/sql"
	"errors"
)

type URLStore struct {
	LongURL  string
	ShortURL string
	Created  string
	Updated  string
	Clicks   int
}

type ShortenerDataModel struct {
	DB *sql.DB
}

// Latest returns the latest URLs
func (m *ShortenerDataModel) Latest() ([]*URLStore, error) {
	stmt := `SELECT original_url, shortened_url, created, updated
	         FROM urls ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := []*URLStore{}

	for rows.Next() {
		url := &URLStore{}

		if err := rows.Scan(&url.LongURL, &url.ShortURL, &url.Created, &url.Updated, url.Clicks); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (m *ShortenerDataModel) GetByShortURL(shortURL string) (*URLStore, error) {
	stmt := `SELECT original_url, shortened_url, clicks FROM urls
	WHERE shortened_url = ?`

	url := &URLStore{}
	err := m.DB.QueryRow(stmt, shortURL).Scan(&url.LongURL, &url.ShortURL, &url.Clicks)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}
	return url, nil
}

func (m *ShortenerDataModel) Insert(original string, shortened string, clicks int) (int64, error) {
	stmt := `INSERT INTO ulrs (original_url, shortened_url, clicks) VALUES(?, ?, ?)`
	result, err := m.DB.Exec(stmt, original, shortened, clicks)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (m *ShortenerDataModel) IncrementClicks(shortURL string) error {
	stmt := `UPDATE urls SET clicks = clicks + 1, updated = CURRENT_TIMESTAMP
	WHERE shortened_url = ?`

	result, err := m.DB.Exec(stmt, shortURL)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("URL not found")
	}

	return nil

}
