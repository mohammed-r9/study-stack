package handler

import (
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/env"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) VerifyEmail(c *fiber.Ctx) error {
	tokenStr := c.Query("t")

	if tokenStr == "" {
		return appErrors.InvalidVerificationToken
	}

	if err := h.svc.VerifyEmail(c.Context(), tokenStr); err != nil {
		return err
	}

	return c.Redirect(env.Config.FRONTEND_URL, fiber.StatusFound)
}
