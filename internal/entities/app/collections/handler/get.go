package handler

import (
	"log"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetCollectionByID(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	collectionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Println(err)
		return appErrors.BadData
	}

	collection, err := h.svc.GetCollectionByID(c.Context(), userData.UserID, collectionID)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(collection)
}

func (h *Handler) GetCollections(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	archived := c.Query("archived")
	var collections []repo.Collection
	var err error

	switch archived {
	case "":
		collections, err = h.svc.GetAllCollections(c.Context(), userData.UserID)
	case "true":
		collections, err = h.svc.GetAllArchived(c.Context(), userData.UserID)
	case "false":
		collections, err = h.svc.GetAllUnarchived(c.Context(), userData.UserID)
	default:
		log.Println("invalid filter in collections")
		return appErrors.BadData
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(collections)
}
