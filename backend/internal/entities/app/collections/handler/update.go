package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type updateReq struct {
	ToArchive   *bool   `json:"to_archive"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (h *Handler) UpdateCollection(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	req := new(updateReq)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	collectionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Printf("error parsing collection id: %v\n", err)
		return appErrors.BadData
	}

	if req.Title != nil {
		if err := h.svc.UpdateTitle(c.Context(), collectionID, userData.UserID, *req.Title); err != nil {
			log.Printf("error updating title: %v\n", err)
			return err
		}
	}

	if req.Description != nil {
		if err := h.svc.UpdateDescription(c.Context(), collectionID, userData.UserID, *req.Description); err != nil {
			log.Printf("error updating description: %v\n", err)
			return err
		}
	}

	if req.ToArchive != nil {
		if err := h.svc.UpdateIsArchived(c.Context(), collectionID, userData.UserID, *req.ToArchive); err != nil {
			log.Printf("error updating archive status: %v\n", err)
			return err
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
