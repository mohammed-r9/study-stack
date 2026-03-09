package handler

import (
	"log"
	"study-stack/internal/entities/app/flashcards/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) DeleteFlashcard(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		log.Println("invalid user data")
		return appErrors.BadData
	}

	flashcardIDStr := c.Params("id")
	if flashcardIDStr == "" {
		log.Println("invalid flashcard id")
		return appErrors.BadData
	}
	flashcardID, err := uuid.Parse(flashcardIDStr)
	if err != nil {
		return err
	}

	deletedID, err := h.svc.DeleteFlashcard(c.Context(), service.DeleteParams{
		UserID:      userData.UserID,
		FlashcardID: flashcardID,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString(deletedID.String())
}
