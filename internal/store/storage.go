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
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		Delete(context.Context, int64) error
		CreateAndInvite(context.Context, *User, string, time.Duration) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Pokemon: &PokemonStore{db: db},
		Users:   &UserStore{db: db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
