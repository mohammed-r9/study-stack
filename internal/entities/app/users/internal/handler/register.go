package handler

import (
	"log"
	"study-stack/internal/entities/app/users/internal/service"
	"study-stack/internal/entities/mailer"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/gofiber/fiber/v2"
)

type registerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *Handler) Register(c *fiber.Ctx) error {
	req := new(registerRequest)

	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("invalid request: %v\n", err)
		return appErrors.BadData
	}

	token, err := h.svc.RegisterUser(c.Context(), service.RegisterParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("error registering user: %v\n", err)
		return err
	}

	if err := mailer.SendVerificationEmail(req.Email, token); err != nil {
		log.Printf("error sending verification email: %v\n", err)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
