package materials

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/materials/internal/handler"
	"study-stack/internal/shared/middleware"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, a *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if a == nil {
			log.Fatalln("Cannot init materials layer with a nil app")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(a, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {
	materials := a.Group("/materials", middleware.Authenticate)

	materials.Post("/", h.InsertMaterial)
	materials.Get("/", h.GetAllMaterials)
	materials.Get("/:id", h.GetMaterialByID)
	materials.Patch("/:id", h.UpdateMaterial)
}
