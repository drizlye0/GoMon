package main

import (
	"net/http"
	"strconv"

	"github.com/drizlye0/GoMon/internal/store"
	"github.com/go-chi/chi/v5"
)

type createPokemonPayload struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Type      []string `json:"type"`
	Region    string   `json:"region"`
	Abilities []string `json:"abilities"`
	Game      string   `json:"game"`
}

func (app *application) getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	value := chi.URLParam(r, "pokemonID")
	id, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	pokemon, err := app.store.Pokemon.GetByID(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, pokemon); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPokemonPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	pokemon := &store.Pokemon{
		ID:        payload.ID,
		Name:      payload.Name,
		Type:      payload.Type,
		Region:    payload.Region,
		Abilities: payload.Abilities,
		Game:      payload.Game,
	}

	if err := app.store.Pokemon.Create(ctx, pokemon); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, pokemon); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
