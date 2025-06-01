package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
