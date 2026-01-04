package users

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/users/internal/handler"
	"sync"

	"github.com/go-chi/chi/v5"
)

var once sync.Once

func Init(db *sql.DB, r *chi.Mux) {
	once.Do(func() {
		if r == nil {
			log.Fatalln("Cannot init users layer with a nil router")
		}
		h := handler.NewHandler(db)
		registerRoutes(r, h)
	})
}

func registerRoutes(r *chi.Mux, h *handler.Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
	})
}
