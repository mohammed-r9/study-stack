package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

func (s *Service) InsertMaterial(ctx context.Context, title string, userID, collectionID uuid.UUID) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	return s.repo.InsertMaterial(ctx, repo.InsertMaterialParams{
		ID:           id,
		CollectionID: collectionID,
		Title:        title,
		UserID:       userID,
	})
}
