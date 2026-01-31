package handler

import (
	"log"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetAllMaterials(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	collectionID, err := uuid.Parse(c.Query("collection_id"))
	if err != nil {
		return appErrors.BadData
	}

	filter := c.Query("archived")

	var materials []repo.Material

	switch filter {
	case "true":
		materials, err = h.svc.GetAllArchived(c.Context(), userData.UserID, collectionID)
	case "false":
		materials, err = h.svc.GetAllUnarchived(c.Context(), userData.UserID, collectionID)
	case "":
		materials, err = h.svc.GetAll(c.Context(), userData.UserID, collectionID)
	default:
		log.Println("invalid filter in materials")
		return appErrors.BadData
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(materials)
}

func (h *Handler) GetMaterialByID(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	idStr := c.Params("id")
	materialID, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Error parsin material id: %v\n", err)
		return appErrors.BadData
	}

	material, err := h.svc.GetByID(c.Context(), userData.UserID, materialID)
	if err != nil {
		log.Printf("Error getting a material: %v\n", err)
		return err
	}

	return c.JSON(material)
}
