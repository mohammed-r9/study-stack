package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

type CreateFlashcardParams struct {
	UserID     uuid.UUID
	MaterialID uuid.UUID
	Front      string
	Back       string
}

func (s *Service) CreateFlashcard(ctx context.Context, params CreateFlashcardParams) error {
	if params.UserID == uuid.Nil {
		return appErrors.BadData
	}

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	return s.repo.CreateFlashcard(ctx, repo.CreateFlashcardParams{
		ID:         id,
		UserID:     params.UserID,
		MaterialID: params.MaterialID,
		Front:      params.Front,
		Back:       params.Back,
	})
}
