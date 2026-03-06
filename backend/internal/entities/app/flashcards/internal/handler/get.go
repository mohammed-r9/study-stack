package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetFlashcards(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	// one | many
	quantity := c.Query("quantity")
	if quantity == "" {
		log.Println("invalid query type")
		return appErrors.BadData
	}

	// to be refactored
	if quantity == "one" {
		flashcard, err := h.svc.GetAndUseFlashcard(c.Context(), userData.UserID)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(flashcard)
	}
	return nil
}
