package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

func (s *Service) GetCollectionByID(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID) (repo.Collection, error) {
	return s.repo.GetCollectionByID(ctx, repo.GetCollectionByIDParams{
		UserID: userID,
		ID:     collectionID,
	})
}

func (s *Service) GetAllArchived(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	return s.repo.GetAllArchivedCollections(ctx, userID)
}

func (s *Service) GetAllUnarchived(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	return s.repo.GetAllUnarchivedCollections(ctx, userID)
}

func (s *Service) GetAllCollections(ctx context.Context, userID uuid.UUID) ([]repo.Collection, error) {
	return s.repo.GetAllCollections(ctx, userID)
}
