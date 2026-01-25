package collections

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/collections/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, a *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if a == nil {
			log.Fatalln("Cannot init collections layer with a nil app")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(a, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {
	collections := a.Group("/collections", middleware.Authenticate)
	collections.Get("/", h.GetCollections)
	collections.Get("/:id", h.GetCollectionByID)
	collections.Post("/", h.CreateCollection)
	collections.Patch("/:id", h.UpdateCollection)
}
