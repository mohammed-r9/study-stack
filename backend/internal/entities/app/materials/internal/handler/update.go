package handler

import (
	"log"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type updateReq struct {
	ToArchive *bool   `json:"to_archive"`
	Title     *string `json:"title"`
}

func (h *Handler) UpdateMaterial(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	req := new(updateReq)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	materialID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Printf("error parsing material id: %v\n", err)
		return appErrors.BadData
	}

	material := repo.Material{}
	if req.Title != nil {
		material, err = h.svc.UpdateMaterialTitle(c.Context(), *req.Title, materialID, userData.UserID)
	}

	if req.ToArchive != nil {
		material, err = h.svc.UpdateMaterialArchivedAt(c.Context(), *req.ToArchive, materialID, userData.UserID)
	}

	if err != nil {
		log.Printf("Error updating material: %v", err)
		return err
	}

	log.Println(material)
	return c.Status(fiber.StatusOK).JSON(material)
}
