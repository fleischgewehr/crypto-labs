package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Client *sql.DB
}

func Get() (*DB, error) {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}
