package users

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/users/internal/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, r *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if r == nil {
			log.Fatalln("Cannot init users layer with a nil router")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(r, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {

	auth := a.Group("/auth")
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/refresh", h.RefreshToken)

	users := a.Group("/users")
	users.Post("/password/reset", h.RequestPasswordReset)
	users.Get("/verify", h.VerifyEmail)

	protected := users.Group("/", middleware.Authenticate)
	protected.Get("/me", h.GetUserByID)
	protected.Patch("/", h.UpdateUser)
}
