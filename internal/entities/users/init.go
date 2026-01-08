package users

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/users/internal/handler"
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
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/refresh", h.RefreshToken)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/password/reset", h.RequestPasswordReset)
		r.Get("/verify", h.VerifyEmail)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Authenticate)
			r.With(middleware.Authenticate).Get("/me", h.GetUserByID)
			r.Patch("/", h.UpdateUser)
		})
	})
}
