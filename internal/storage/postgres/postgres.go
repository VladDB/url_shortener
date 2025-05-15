package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	const op = "storage.sql.NewStorage"

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname='storage')").Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !exists {
		// Если базы нет, создаём её
		_, err = db.Exec("CREATE DATABASE storage")
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	storageDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = storageDB.Exec(`
	CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias on url(alias);
	`)

	if err != nil {
		storageDB.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: storageDB}, nil
}
