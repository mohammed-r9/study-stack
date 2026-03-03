package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetAndUseFlashcard(ctx context.Context, userID uuid.UUID) (repo.Flashcard, error) {
	if userID == uuid.Nil {
		return repo.Flashcard{}, appErrors.BadData
	}

	return s.repo.GetAndUseFlashCard(ctx, s.db, userID)
}
