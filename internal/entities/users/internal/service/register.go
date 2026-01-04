package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/utils"

	"github.com/google/uuid"
)

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

func (s *Service) RegisterUser(ctx context.Context, params RegisterParams) error {
	password := utils.Password{}
	err := password.Set(params.Password)
	if err != nil {
		return err
	}

	err = s.repo.InsertUser(ctx, repo.InsertUserParams{
		ID:           uuid.New(),
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: password.Hash,
		Salt:         password.Salt,
	})

	return err
}
