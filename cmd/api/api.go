package main

import (
	"log"
	"net/http"
	"time"

	"github.com/drizlye0/GoMon/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr               string
	maxIdleConnections int
	maxOpenConnections int
	maxIdleTime        string
}

type application struct {
	config config
	store  *store.Storage
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.checkHealthHandler)

		r.Route("/pokemon", func(r chi.Router) {
			r.Post("/", app.createPokemonHandler)

			r.Route("/{pokemonID}", func(r chi.Router) {
				r.Get("/", app.getPokemonHandler)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	server := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server listening on port %s", app.config.addr)
	return server.ListenAndServe()
}
