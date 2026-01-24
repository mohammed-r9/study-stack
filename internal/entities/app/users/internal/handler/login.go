package handler

import (
	"log"
	"study-stack/internal/entities/app/users/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(loginRequest)

	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("invalid request: %v\n", err)
		return appErrors.BadData
	}

	tokens, err := h.svc.Login(c.Context(), service.LoginParams{
		Email:       req.Email,
		Password:    req.Password,
		Device_name: utils.GetDeviceNameFromUserAgent(c.Get("User-Agent")),
	})
	if err != nil {
		return err
	}

	utils.SetRefreshCookieFiber(c, tokens.Refresh)
	utils.SetCsrfCookieFiber(c, tokens.Csrf)

	return c.JSON(fiber.Map{
		"access_token": tokens.Access,
	})
}
