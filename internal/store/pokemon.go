package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Pokemon struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Type       []string `json:"type"`
	Region     string   `json:"region"`
	Abilities  []string `json:"abilities"`
	Game       []string `json:"game"`
	Created_At string   `json:"created_at"`
	Updated_At string   `json:"updated_at"`
}

type PokemonStore struct {
	db *sql.DB
}

func (s *PokemonStore) Create(ctx context.Context, pokemon *Pokemon) error {
	query := `
		INSERT INTO pokemons(id, name, type, region, abilities, game)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		pokemon.ID,
		pokemon.Name,
		pq.Array(pokemon.Type),
		pokemon.Region,
		pq.Array(pokemon.Abilities),
		pq.Array(pokemon.Game),
	).Scan(
		&pokemon.Created_At,
		&pokemon.Updated_At,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PokemonStore) GetByID(ctx context.Context, id int64) (*Pokemon, error) {
	query := `
		SELECT 
			id, name, type, region, game, abilities, created_at, updated_at
		FROM pokemons
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	var pokemon Pokemon

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&pokemon.ID,
		&pokemon.Name,
		pq.Array(&pokemon.Type),
		&pokemon.Region,
		pq.Array(&pokemon.Game),
		pq.Array(&pokemon.Abilities),
		&pokemon.Created_At,
		&pokemon.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
