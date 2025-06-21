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
		Create(context.Context, *Pokemon) error
		GetByID(context.Context, int64) (*Pokemon, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
		Delete(context.Context, int64) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Pokemon: &PokemonStore{db: db},
		Users: &UserStore{db}
	}
}
