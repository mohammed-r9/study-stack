package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetCollectionByID(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID) (repo.Collection, error) {
	if userID == uuid.Nil {
		return repo.Collection{}, appErrors.InvalidUserID
	}
	return s.repo.GetCollectionByID(ctx, repo.GetCollectionByIDParams{
		UserID: userID,
		ID:     collectionID,
	})
}

