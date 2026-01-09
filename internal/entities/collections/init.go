package collections

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/collections/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var once sync.Once

func Init(db *sql.DB, r *chi.Mux, v *validator.Validate) {
	once.Do(func() {
		if r == nil {
			log.Fatalln("Cannot init users layer with a nil router")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(r, h)
	})
}

func registerRoutes(r *chi.Mux, h *handler.Handler) {
	r.Route("/collections", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/{id}", h.GetCollectionByID)
		r.Get("/", h.GetAllCollections)
		r.Get("/archived", h.GetAllArchivedCollections)
		r.Get("/unarchived", h.GetAllUnarchivedCollections)
		r.Post("/", h.CreateCollection)
		r.Patch("/{id}", h.UpdateCollection)
	})
}
