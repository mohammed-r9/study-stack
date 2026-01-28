package service

import (
	"context"
	"database/sql"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetCollectionByID(ctx context.Context, userID, collectionID uuid.UUID) (repo.Collection, error) {
	collection, err := s.repo.GetCollectionByID(ctx, repo.GetCollectionByIDParams{
		UserID: userID,
		ID:     collectionID,
	})
	if err == sql.ErrNoRows {
		return repo.Collection{}, appErrors.NotFound
	}
	return collection, err
}

func (s *Service) GetAllArchived(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	collections, err := s.repo.GetAllArchivedCollections(ctx, userID)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (s *Service) GetAllUnarchived(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	collections, err := s.repo.GetAllUnarchivedCollections(ctx, userID)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (s *Service) GetAllCollections(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	collections, err := s.repo.GetAllCollections(ctx, userID)
	if err != nil {
		return nil, err
	}
	return collections, nil
}
