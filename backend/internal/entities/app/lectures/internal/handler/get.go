package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetAll(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		log.Println("invalid user")
		return appErrors.BadData
	}

	lastSeenIdStr := c.Query("lastSeenIdStr", "")
	lastSeenId, err := utils.ParseOptionalUUID(lastSeenIdStr)
	if err != nil {
		return err
	}

	lectures, err := h.svc.GetAllLectures(c.Context(), userData.UserID, lastSeenId)
	if err != nil {
		return err
	}
	return c.JSON(lectures)
}
