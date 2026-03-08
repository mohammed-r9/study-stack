package handler

import (
	"log"
	"study-stack/internal/entities/app/flashcards/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetFlashcards(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadData
	}

	// study | list
	quantity := c.Query("mode")
	if quantity == "" {
		log.Println("invalid query type")
		return appErrors.BadData
	}

	switch quantity {
	case "study":
		flashcard, err := h.svc.GetAndUseFlashcard(c.Context(), userData.UserID)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(flashcard)

	case "list":
		lastSeenIdStr := c.Query("last_seen_flashcard_id", "")
		lastSeenId, err := utils.ParseOptionalUUID(lastSeenIdStr)
		if err != nil {
			return err
		}
		flashcardsPage, err := h.svc.GetFlashcardPage(c.Context(), service.GetFlashcardPageParams{
			UserID:            userData.UserID,
			LastSeenFlashcard: lastSeenId,
		})
		if err != nil {
			return err
		}
		return c.JSON(flashcardsPage)
	}
	return appErrors.BadData
}
