package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

type DeleteParams struct {
	UserID      uuid.UUID
	FlashcardID uuid.UUID
}

func (s *Service) DeleteFlashcard(ctx context.Context, params DeleteParams) error {
	if params.UserID == uuid.Nil || params.FlashcardID == uuid.Nil {
		return appErrors.BadData
	}

	rowsAffected, err := s.repo.DeleteFlashcard(ctx, repo.DeleteFlashcardParams{
		FlashcardID: params.FlashcardID,
		UserID:      params.UserID,
	})

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return appErrors.NotFound
	}

	return nil
}
