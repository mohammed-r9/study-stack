package handler

import (
	"log"
	"study-stack/internal/entities/app/flashcards/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

func (h *Handler) UpdateFlashcard(c *fiber.Ctx) error {
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

	req := new(createRequest)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadData
	}

	if req.Back == "" && req.Front == "" {
		return appErrors.BadData
	}

	flashcard, err := h.svc.UpdateFlashcard(c.Context(), service.UpdateFlashcardParams{
		UserID:      userData.UserID,
		FlashcardID: flashcardID,
		Front:       req.Front,
		Back:        req.Back,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(flashcard)
}
