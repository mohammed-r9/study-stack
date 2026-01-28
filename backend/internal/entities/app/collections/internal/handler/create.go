package handler

import (
	"log"
	"study-stack/internal/entities/app/collections/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

type createReq struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *Handler) CreateCollection(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	req := new(createReq)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("invalid request: %v\n", err)
		return appErrors.BadData
	}

	if err := h.svc.CreateCollection(c.Context(), service.CreateCollectionParams{
		UserID:     userData.UserID,
		Title:      req.Title,
		Desription: req.Description,
	}); err != nil {
		log.Printf("error while creating collection: %v\n", err)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
