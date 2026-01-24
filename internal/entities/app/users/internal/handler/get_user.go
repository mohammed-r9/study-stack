package handler

import (
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	user, err := h.svc.GetUserByID(c.Context(), userData.UserID)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
