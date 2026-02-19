package handler

import (
	"log"
	"study-stack/internal/entities/app/lectures/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetSignedURL(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		log.Println("invalid user")
		return appErrors.BadData
	}

	lectureIDStr := c.Params("id", "")
	if lectureIDStr == "" {
		return appErrors.BadData
	}
	lectureID, err := uuid.Parse(lectureIDStr)
	if err != nil {
		return err
	}

	url, err := h.svc.GetSignedURL(c.Context(), service.GetSignedURLParams{UserID: userData.UserID, LectureID: lectureID})

	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"url": url,
	})
}
