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

func Init(db *sql.DB, app *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if app == nil {
			log.Fatalln("Cannot init materials layer with a nil app")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(app, h)
	})
}

func registerRoutes(app *fiber.App, h *handler.Handler) {
	materials := app.Group("/materials", middleware.Authenticate)

	materials.Post("/", h.InsertMaterial)
	materials.Get("/", h.GetAllMaterials)
	materials.Get("/:id", h.GetMaterialByID)
	materials.Patch("/:id", h.UpdateMaterial)
}
