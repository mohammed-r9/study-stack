package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type materialCreationReq struct {
	CollectionID uuid.UUID `json:"collection_id" validate:"required"`
	Title        string    `json:"title" validate:"required,min=4"`
}

func (h *Handler) InsertMaterial(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	req := new(materialCreationReq)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("error decoding stuct: %v\n", err)
		return appErrors.BadData
	}

	if err := h.svc.InsertMaterial(c.Context(), req.Title, userData.UserID, req.CollectionID); err != nil {
		log.Printf("error inseting material: %v\n", err)
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
