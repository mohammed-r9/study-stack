package utils

import (
	"github.com/gofiber/fiber/v2"
)

func BuildRatelimitKeyForAuthedUser(c *fiber.Ctx, prefix string) string {
	userData, _ := DataFromLocals(c)
	return prefix + ":" + userData.UserID.String()
}

func BuildRatelimitKeyForRefresh(c *fiber.Ctx, prefix string) string {
	refreshToken := c.Cookies("refresh_token")
	return prefix + ":" + refreshToken
}

func BuildRatelimitKeyForPublicRoutes(c *fiber.Ctx, prefix string) string {
	ip := c.IP()
	return prefix + ":" + ip
}
