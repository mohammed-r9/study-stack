package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUserLibrary(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	lib, err := h.svc.GetUserLibrary(c.Context(), userData.UserID)
	if err != nil {
		log.Printf("Error while getting library: %v", err)
		return err
	}

	return c.JSON(lib)
}
