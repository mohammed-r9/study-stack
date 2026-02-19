package handler

import (
	"log"
	"study-stack/internal/entities/app/lectures/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetAll(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		log.Println("invalid user")
		return appErrors.BadData
	}

	lastSeenIdStr := c.Query("last_seen", "")
	lastSeenId, err := utils.ParseOptionalUUID(lastSeenIdStr)
	if err != nil {
		return err
	}
	materialIdStr := c.Query("m_id", "")
	if materialIdStr == "" {
		return appErrors.BadData
	}
	materialID, err := uuid.Parse(materialIdStr)
	if err != nil {
		return err
	}
	lectures, err := h.svc.GetAllLectures(c.Context(), service.GetAllLecturesParams{
		UserID:     userData.UserID,
		MaterialID: materialID,
		LastSeenID: lastSeenId,
	})
	if err != nil {
		return err
	}
	return c.JSON(lectures)
}
