package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) UpdateIsArchived(ctx context.Context, collectionID, userID uuid.UUID, val bool) error {
	if val == true {
		rowsAffected, err := s.repo.ArchiveCollection(ctx, repo.ArchiveCollectionParams{
			ID:     collectionID,
			UserID: userID,
		})

		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return appErrors.NoRowsAffected
		}
	}

	rowsAffected, err := s.repo.UnarchiveCollection(ctx, repo.UnarchiveCollectionParams{
		ID:     collectionID,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return appErrors.NoRowsAffected
	}

	return nil
}

func (s *Service) UpdateTitle(ctx context.Context, collectionID, userID uuid.UUID, title string) error {
	if len(title) < 3 {
		return appErrors.InvalidData
	}

	rowsAffected, err := s.repo.UpdateCollectionTitle(ctx, repo.UpdateCollectionTitleParams{
		Title:  title,
		ID:     collectionID,
		UserID: userID,
	})

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return appErrors.NoRowsAffected
	}
	return nil

}

func (s *Service) UpdateDescription(ctx context.Context, collectionID, userID uuid.UUID, desc string) error {
	if len(desc) < 3 {
		return appErrors.InvalidData
	}

	rowsAffected, err := s.repo.UpdateCollectionDescription(ctx, repo.UpdateCollectionDescriptionParams{
		Description: desc,
		ID:          collectionID,
		UserID:      userID,
	})

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return appErrors.NoRowsAffected
	}
	return nil

}
