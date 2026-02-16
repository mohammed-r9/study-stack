package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/entities/tokens/stateful"
	"study-stack/internal/shared/password"
	"time"

	"github.com/google/uuid"
)

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

// returns a verifiction token on success
func (s *Service) RegisterUser(ctx context.Context, params RegisterParams) (string, error) {
	password := password.Password{}
	err := password.Set(params.Password)
	if err != nil {
		return "", err
	}

	userID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	err = s.repo.InsertUser(ctx, repo.InsertUserParams{
		ID:           userID,
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: password.Hash,
		Salt:         password.Salt,
	})

	if err != nil {
		return "", err
	}

	token, err := stateful.NewOpaqueToken(userID, stateful.EmailVerification)
	if err != nil {
		return "", err
	}

	err = s.repo.InsertToken(ctx, repo.InsertTokenParams{
		UserID:    userID,
		Hash:      token.Hash,
		Scope:     string(token.Scope),
		ExpiresAt: time.Now().Add(stateful.EmailVerificationTTL),
	})
	if err != nil {
		return "", err
	}

	return token.PlainText, nil
}
