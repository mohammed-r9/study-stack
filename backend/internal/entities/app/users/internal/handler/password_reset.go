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
		return appErrors.BadData
	}

	token, err := h.svc.RequestPasswordReset(c.Context(), req.Email)
	if err != nil {
		log.Printf("error generating password reset token: %v\n", err)
		return err
	}

	_ = token // send email later

	return c.SendStatus(fiber.StatusOK)
}

type confirmReq struct {
	NewPassword string `json:"new_password"`
	Token       string `json:"token"`
}

func (h *Handler) ConfirmPasswordReset(c *fiber.Ctx) error {
	req := new(confirmReq)

	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	err := h.svc.ConfirmPasswordReset(c.Context(), req.Token, req.NewPassword)
	if err != nil {
		log.Printf("error resetting password: %v\n", err)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
