package handler

import (
	"log"

	appErrors "study-stack/internal/shared/app_errors"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		log.Println("no refresh cookie present")
		return appErrors.Unauthorized
	}

	csrf := c.Get("X-CSRF-Token")
	if csrf == "" {
		log.Println("no CSRF token present")
		return appErrors.Unauthorized
	}

	accessToken, err := h.svc.RefreshToken(c.Context(), refreshToken, csrf)
	if err != nil {
		log.Printf("error generating the jwt: %v\n", err)
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": accessToken,
	})
}
