package store

import (
	"context"
	"database/sql"
	"time"
)

var (
	QueryContextTimeout = time.Second * 5
)

type Storage struct {
	Pokemon interface {
		GetByID(context.Context, int64) (*Pokemon, error)
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Pokemon: &PokemonStore{db: db},
	}
}
