package service

import (
	"context"
	"database/sql"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetByID(ctx context.Context, userID, materialID uuid.UUID) (repo.Material, error) {
	material, err := s.repo.GetMaterialByID(ctx, repo.GetMaterialByIDParams{
		UserID: userID,
		ID:     materialID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return repo.Material{}, appErrors.NotFound
		}
		return repo.Material{}, err
	}

	return material, nil
}

func (s *Service) GetAllUnarchived(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	materials, err := s.repo.GetAllUnarchivedMaterialsInCollection(ctx, repo.GetAllUnarchivedMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})

	if err != nil {
		return nil, err
	}
	if len(materials) == 0 {
		return nil, appErrors.NotFound
	}
	return materials, nil
}

func (s *Service) GetAllArchived(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	materials, err := s.repo.GetAllArchivedMaterialsInCollection(ctx, repo.GetAllArchivedMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})

	if err != nil {
		return nil, err
	}
	if len(materials) == 0 {
		return nil, appErrors.NotFound
	}
	return materials, nil
}

func (s *Service) GetAll(ctx context.Context, userID, collectionID uuid.UUID) ([]repo.Material, error) {
	materials, err := s.repo.GetAllMaterialsInCollection(ctx, repo.GetAllMaterialsInCollectionParams{
		UserID:       userID,
		CollectionID: collectionID,
	})

	if err != nil {
		return nil, err
	}
	if len(materials) == 0 {
		return nil, appErrors.NotFound
	}
	return materials, nil
}
