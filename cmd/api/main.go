package main

import (
	"log"

	"github.com/drizlye0/GoMon/internal/db"
	"github.com/drizlye0/GoMon/internal/env"
	"github.com/drizlye0/GoMon/internal/store"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8000"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://user:userpass@localhost/GoMon?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConnections,
		cfg.db.maxIdleConnections,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store := store.NewStorage(db)

	app := application{
		config: *cfg,
		store:  store,
	}

	router := app.mount()
	app.run(router)
}
