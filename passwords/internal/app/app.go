package app

import (
	"github.com/fleischgewehr/crypto-labs/passwords/internal/db"
)

type Application struct {
	DB *db.DB
}

func Get() (*Application, error) {
	db, err := db.Get()
	if err != nil {
		return nil, err
	}

	return &Application{
		DB: db,
	}, nil
}
