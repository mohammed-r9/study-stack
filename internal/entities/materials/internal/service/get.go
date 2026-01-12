package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

func (s *Service) GetByID(ctx context.Context, userID, materialID uuid.UUID) (repo.Material, error) {
	return s.repo.GetMaterialByID(ctx, repo.GetMaterialByIDParams{
		UserID: userID,
		ID:     materialID,
	})
}

func (s *Service) GetAllUnarchived(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	return s.repo.GetAllUnarchivedMaterialsInCollection(ctx, repo.GetAllUnarchivedMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
}

func (s *Service) GetAllArchived(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	return s.repo.GetAllArchivedMaterialsInCollection(ctx, repo.GetAllArchivedMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
}

func (s *Service) GetAll(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	return s.repo.GetAllMaterialsInCollection(ctx, repo.GetAllMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
}
