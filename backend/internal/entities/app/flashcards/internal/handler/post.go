package handler

import (
	"log"
	"study-stack/internal/entities/app/flashcards/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type createRequest struct {
	Front      string    `json:"front" validate:"required,min=4"`
	Back       string    `json:"back" validate:"required,min=4"`
	MaterialID uuid.UUID `json:"material_id" validate:"required"`
}

func (h *Handler) CreateFlashcard(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	req := new(createRequest)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("error decoding stuct: %v\n", err)
		return appErrors.BadData
	}

	err := h.svc.CreateFlashcard(c.Context(), service.CreateFlashcardParams{
		UserID:     userData.UserID,
		MaterialID: req.MaterialID,
		Front:      req.Front,
		Back:       req.Back,
	})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
