package main

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/collections"
	"study-stack/internal/entities/app/materials"
	"study-stack/internal/entities/app/users"
	"study-stack/internal/entities/mailer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type application struct {
	db     *sql.DB
	router *fiber.App
	addr   string
}

func (a *application) mount() {

	a.router.Use(requestid.New())
	a.router.Use(logger.New())
	a.router.Use(recover.New())

	a.router.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status Is Available")
	})

	validator := validator.New()
	mailer.Init()
	users.Init(a.db, a.router, validator)
	collections.Init(a.db, a.router, validator)
	materials.Init(a.db, a.router, validator)
}

func (a *application) run() error {

	log.Printf("Server Has Started At Addr %s", a.addr)

	return a.router.Listen(a.addr)
}
