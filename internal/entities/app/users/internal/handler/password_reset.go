package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	Email string `json:"email"`
}

func (h *Handler) RequestPasswordReset(c *fiber.Ctx) error {
	req := new(request)

	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadRequest
	}

	token, err := h.svc.RequestPasswordReset(c.Context(), req.Email)
	if err != nil {
		log.Printf("error generating password reset token: %v\n", err)
		return err
	}

	_ = token // send email later

	return c.SendStatus(fiber.StatusOK)
}
