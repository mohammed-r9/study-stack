package users

import (
	"database/sql"
	"log"
	"study-stack/internal/entities/app/users/internal/handler"
	"study-stack/internal/shared/consts"
	"study-stack/internal/shared/middleware"
	"study-stack/internal/shared/utils"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Init(db *sql.DB, a *fiber.App, v *validator.Validate) {
	once.Do(func() {
		if a == nil {
			log.Fatalln("Cannot init users layer with a nil router")
		}
		h := handler.NewHandler(db, v)
		registerRoutes(a, h)
	})
}

func registerRoutes(a *fiber.App, h *handler.Handler) {
	ipLimiter := middleware.RateLimitMiddleware(middleware.RateLimitConfig{
		Max:        10,
		Window:     time.Minute,
		KeyBuilder: utils.BuildRatelimitKeyForPublicRoutes,
		Prefix:     consts.RL_PUBLIC,
	})

	refreshLimiter := middleware.RateLimitMiddleware(middleware.RateLimitConfig{
		Max:        10,
		Window:     time.Minute,
		KeyBuilder: utils.BuildRatelimitKeyForPublicRoutes,
		Prefix:     consts.RL_REFRESH,
	})

	authLimiter := middleware.RateLimitMiddleware(middleware.RateLimitConfig{
		Max:        10,
		Window:     time.Minute,
		KeyBuilder: utils.BuildRatelimitKeyForPublicRoutes,
		Prefix:     consts.RL_AUTH,
	})

	users := a.Group("/users", ipLimiter)
	users.Post("/password/reset", h.RequestPasswordReset)
	users.Post("/password/reset/confirm", h.ConfirmPasswordReset)
	users.Get("/verify", h.VerifyEmail)
	users.Post("/register", h.Register)
	users.Post("/login", h.Login)
	users.Post("/refresh", refreshLimiter, h.RefreshToken)

	protected := users.Group("/", middleware.Authenticate, authLimiter)
	protected.Get("/me", h.GetUserByID)
	protected.Patch("/", h.UpdateUser)
	protected.Get("/me/library", h.GetUserLibrary)
}
