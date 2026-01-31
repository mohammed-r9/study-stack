package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

func (s *Service) InsertMaterial(ctx context.Context, title string, userID, collectionID uuid.UUID) error {
	return s.repo.InsertMaterial(ctx, repo.InsertMaterialParams{
		ID:           uuid.New(),
		CollectionID: collectionID,
		Title:        title,
		UserID:       userID,
	})
}
