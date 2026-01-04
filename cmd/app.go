package main

import (
	"database/sql"
	"log"
	"net/http"
	"study-stack/internal/entities/users"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	db     *sql.DB
	router *chi.Mux
	addr   string
}

func (a *application) mount() {
	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	a.router.Use(middleware.Timeout(60 * time.Second))

	a.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Status Is Available"))
	})
	users.Init(a.db, a.router)
}

func (a *application) run() error {

	server := &http.Server{
		Addr:         a.addr,
		Handler:      a.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server Has Started At Addr %s", a.addr)
	return server.ListenAndServe()
}
