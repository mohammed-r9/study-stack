package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (repo.GetUserByIDRow, error) {
	if userID == uuid.Nil {
		return repo.GetUserByIDRow{}, appErrors.BadRequest
	}
	return s.repo.GetUserByID(ctx, userID)

}
