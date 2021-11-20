package app

import (
	"github.com/fleischgewehr/crypto-labs/passwords/internal/cache"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/db"
)

type Application struct {
	DB    *db.DB
	Cache *cache.Cache
}

func Get() (*Application, error) {
	db, err := db.Get()
	if err != nil {
		return nil, err
	}
	cache, err := cache.Get()
	if err != nil {
		return nil, err
	}

	return &Application{
		DB:    db,
		Cache: cache,
	}, nil
}
