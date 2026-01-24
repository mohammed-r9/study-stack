package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) UpdateMaterialTitle(ctx context.Context, newTitle string, materialID, userID uuid.UUID) error {
	rowsAffected, err := s.repo.UpdateMaterialTitle(ctx, repo.UpdateMaterialTitleParams{
		Title:  newTitle,
		ID:     materialID,
		UserID: userID,
	})
	if rowsAffected == 0 {
		return appErrors.NotFound
	}
	return err
}

func (s *Service) UpdateMaterialArchivedAt(ctx context.Context, val bool, materialID, userID uuid.UUID) error {
	var rowsAffected int64
	var err error
	if val {
		rowsAffected, err = s.repo.ArchiveMaterial(ctx, repo.ArchiveMaterialParams{
			ID:     materialID,
			UserID: userID,
		})

		if rowsAffected == 0 {
			return appErrors.NotFound
		}
		return err
	}

	rowsAffected, err = s.repo.UnarchiveMaterial(ctx, repo.UnarchiveMaterialParams{
		ID:     materialID,
		UserID: userID,
	})

	if rowsAffected == 0 {
		return appErrors.NotFound
	}
	return err

}
