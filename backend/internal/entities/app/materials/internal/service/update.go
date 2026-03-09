package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

func (s *Service) UpdateMaterialTitle(ctx context.Context, newTitle string, materialID, userID uuid.UUID) (repo.Material, error) {
	return s.repo.UpdateMaterialTitle(ctx, repo.UpdateMaterialTitleParams{
		Title:  newTitle,
		ID:     materialID,
		UserID: userID,
	})
}

func (s *Service) UpdateMaterialArchivedAt(ctx context.Context, val bool, materialID, userID uuid.UUID) (repo.Material, error) {
	if val {
		return s.repo.ArchiveMaterial(ctx, repo.ArchiveMaterialParams{
			ID:     materialID,
			UserID: userID,
		})

	}

	return s.repo.UnarchiveMaterial(ctx, repo.UnarchiveMaterialParams{
		ID:     materialID,
		UserID: userID,
	})

}
