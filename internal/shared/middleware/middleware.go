package middleware

import (
	"log"
	"strings"
	"study-stack/internal/entities/tokens/stateless"
	"study-stack/internal/shared/consts"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	const prefix = "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := strings.TrimPrefix(authHeader, prefix)

	userData, err := stateless.VerifyAccessToken(token)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals(consts.UserDataContextKey, userData)

	return c.Next()
}
