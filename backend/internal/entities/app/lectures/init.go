package lectures

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/lectures/internal/handler"
	S3 "study-stack/internal/s3"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, a *fiber.App, v *validator.Validate, b S3.Storage) {
	once.Do(func() {
		if a == nil {
			log.Fatalln("Cannot init users lectures with a nil router")
		}
		h := handler.NewHandler(db, v, b)
		registerRoutes(a, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {
	// TODO: Define routes
}
