package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

type CreateCollectionParams struct {
	UserID      uuid.UUID
	Title       string
	Description string
}

func (s *Service) CreateCollection(ctx context.Context, params CreateCollectionParams) (repo.Collection, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return repo.Collection{}, err
	}
	count, err := s.repo.GetCollectionsCount(ctx, params.UserID)
	if err != nil {
		return repo.Collection{}, err
	}

	if count >= 20 {
		return repo.Collection{}, appErrors.Forbidden
	}
	return s.repo.CreateCollection(ctx, repo.CreateCollectionParams{
		ID:          id,
		UserID:      params.UserID,
		Title:       params.Title,
		Description: params.Description,
	})
}
