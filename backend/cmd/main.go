package main

import (
	"flag"
	"fmt"
	"log"
	"study-stack/internal/db"
	"study-stack/internal/shared/env"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Sets The Server Port")
	flag.Parse()

	env.LoadEnv("./.env")
	app := application{
		db:   db.NewDb(),
		addr: fmt.Sprintf(":%d", port),
		router: fiber.New(fiber.Config{
			ErrorHandler: FiberErrorHandler,
			BodyLimit:    50 * 1024 * 1024,
		}),
	}

	app.mount()
	if err := app.run(); err != nil {
		log.Fatalf("Failed To Run Application: %v", err)
	}
}
