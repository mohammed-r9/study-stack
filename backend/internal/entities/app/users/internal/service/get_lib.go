package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]repo.GetUserLibraryRow, error) {
	if userID == uuid.Nil {
		return []repo.GetUserLibraryRow{}, appErrors.BadData
	}
	return s.repo.GetUserLibrary(ctx, userID)
}
