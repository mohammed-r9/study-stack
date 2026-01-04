package main

import (
	"flag"
	"fmt"
	"log"
	"study-stack/internal/db"
	"study-stack/internal/env"

	"github.com/go-chi/chi/v5"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Sets The Server Port")
	flag.Parse()

	env.LoadEnv("./.env")
	app := application{
		db:     db.NewDb(),
		addr:   fmt.Sprintf(":%d", port),
		router: chi.NewMux(),
	}

	app.mount()
	if err := app.run(); err != nil {
		log.Fatalf("Failed To Run Application: %v", err)
	}
}
