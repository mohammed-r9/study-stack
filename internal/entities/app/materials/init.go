package materials

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/materials/internal/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var once sync.Once

func Init(db *sql.DB, r *chi.Mux, v *validator.Validate) {
	once.Do(func() {
		if r == nil {
			log.Fatalln("Cannot init materials layer with a nil router")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(r, h)
	})
}

func registerRoutes(r *chi.Mux, h *handler.Handler) {
	r.Route("/materials", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Post("/", h.InsertMaterial)
		r.Get("/", h.GetAllMaterials)
		r.Get("/{id}", h.GetMaterialByID)
		r.Patch("/{id}", h.UpdateMaterial)
	})
}
