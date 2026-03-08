package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

type UpdateFlashcardParams struct {
	UserID      uuid.UUID
	FlashcardID uuid.UUID
	Front       string
	Back        string
}

func (s *Service) UpdateFlashcard(ctx context.Context, params UpdateFlashcardParams) error {
	if (params.UserID == uuid.Nil || params.FlashcardID == uuid.Nil) || (params.Front == "" && params.Back == "") {
		return appErrors.BadData
	}

	rowsAffected, err := s.repo.UpdateFlashcard(ctx, repo.UpdateFlashcardParams{
		UserID:      params.UserID,
		FlashcardID: params.FlashcardID,
		Front:       params.Front,
		Back:        params.Back,
	})

	if rowsAffected == 0 {
		return appErrors.NoChange
	}

	return err
}
