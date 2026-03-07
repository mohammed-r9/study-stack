package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type KeyFunc func(c *fiber.Ctx, prefix string) string

type RateLimitConfig struct {
	Max        int
	Window     time.Duration
	KeyBuilder KeyFunc
	Message    string
	// a prefix for the key
	Prefix string
}

func RateLimitMiddleware(cfg RateLimitConfig) fiber.Handler {
	if cfg.Message == "" {
		cfg.Message = "Too many requests"
	}

	return limiter.New(limiter.Config{
		Max:        cfg.Max,
		Expiration: cfg.Window,
		KeyGenerator: func(c *fiber.Ctx) string {
			return cfg.KeyBuilder(c, cfg.Prefix)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": cfg.Message,
			})
		},
	})
}
