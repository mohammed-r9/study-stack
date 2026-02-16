package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

type CreateCollectionParams struct {
	UserID     uuid.UUID
	Title      string
	Desription string
}

func (s *Service) CreateCollection(ctx context.Context, params CreateCollectionParams) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	return s.repo.CreateCollection(ctx, repo.CreateCollectionParams{
		ID:          id,
		UserID:      params.UserID,
		Title:       params.Title,
		Description: params.Desription,
	})
}
