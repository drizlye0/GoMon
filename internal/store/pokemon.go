package store

import (
	"context"
	"database/sql"
)

type Pokemon struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Type       []string `json:"type"`
	Region     string   `json:"region"`
	Abilites   string   `json:"abilities"`
	Game       string   `json:"game"`
	Created_At string   `json:"created_at"`
	Updated_At string   `json:"updated_at"`
}

type PokemonStore struct {
	db *sql.DB
}

func (s *PokemonStore) GetByID(ctx context.Context, id int64) (*Pokemon, error) {
	query := `
		SELECT 
			name, type, region, game, abilities, created_at, updated_at
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
		&pokemon.Name,
		&pokemon.Type,
		&pokemon.Region,
		&pokemon.Game,
		&pokemon.Abilites,
		&pokemon.Created_At,
		&pokemon.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
