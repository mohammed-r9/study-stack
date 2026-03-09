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

func (s *Service) DeleteFlashcard(ctx context.Context, params DeleteParams) (uuid.UUID, error) {
	if params.UserID == uuid.Nil || params.FlashcardID == uuid.Nil {
		return uuid.Nil, appErrors.BadData
	}

	return s.repo.DeleteFlashcard(ctx, repo.DeleteFlashcardParams{
		FlashcardID: params.FlashcardID,
		UserID:      params.UserID,
	})

}
