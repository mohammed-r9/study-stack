package flashcards

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/flashcards/internal/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, a *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if a == nil {
			log.Fatalln("Cannot init users flashcards with a nil router")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(a, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {
	flashcards := a.Group("/flashcards", middleware.Authenticate)
	flashcards.Post("/", h.CreateFlashcard)
	flashcards.Get("/", h.GetFlashcards)
}
