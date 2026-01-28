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
	err := s.repo.CreateCollection(ctx, repo.CreateCollectionParams{
		ID:          uuid.New(),
		UserID:      params.UserID,
		Title:       params.Title,
		Description: params.Desription,
	})
	if err != nil {
		return err
	}

	return nil
}
