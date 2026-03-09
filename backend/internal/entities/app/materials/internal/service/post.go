package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) InsertMaterial(ctx context.Context, title string, userID, collectionID uuid.UUID) (repo.Material, error) {
	count, err := s.repo.GetMaterialsCount(ctx, repo.GetMaterialsCountParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
	if err != nil {
		return repo.Material{}, err
	}
	if count >= 20 {
		return repo.Material{}, appErrors.Forbidden
	}

	id, err := uuid.NewV7()
	if err != nil {
		return repo.Material{}, err
	}
	return s.repo.InsertMaterial(ctx, repo.InsertMaterialParams{
		ID:           id,
		CollectionID: collectionID,
		Title:        title,
		UserID:       userID,
	})
}
