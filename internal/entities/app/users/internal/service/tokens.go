package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/entities/tokens/stateful"
	appErrors "study-stack/internal/shared/app_errors"
	"time"

	"github.com/google/uuid"
)

func (s *Service) RequestPasswordReset(ctx context.Context, email string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	token, err := stateful.NewOpaqueToken(user.ID, stateful.PasswordReset)
	if err != nil {
		return "", err
	}

	err = s.repo.InsertToken(ctx, repo.InsertTokenParams{
		UserID:    token.UserID,
		Hash:      token.Hash,
		Scope:     string(token.Scope),
		ExpiresAt: time.Now().Add(stateful.PasswordResetTTL),
	})
	if err != nil {
		return "", err
	}

	return token.PlainText, nil
}

func (s *Service) NewEmailVerificationToken(ctx context.Context, userID uuid.UUID) (string, error) {
	if userID == uuid.Nil {
		return "", appErrors.BadRequest
	}

	token, err := stateful.NewOpaqueToken(userID, stateful.EmailVerification)
	if err != nil {
		return "", err
	}

	err = s.repo.InsertToken(ctx, repo.InsertTokenParams{
		UserID:    token.UserID,
		Hash:      token.Hash,
		Scope:     string(token.Scope),
		ExpiresAt: time.Now().Add(stateful.PasswordResetTTL),
	})
	if err != nil {
		return "", err
	}

	return token.PlainText, nil
}
