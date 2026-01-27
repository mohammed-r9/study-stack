package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/entities/tokens/stateful"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/password"
	"time"
)

func (s *Service) ConfirmPasswordReset(ctx context.Context, token string, newPassword string) error {
	tokenHash := stateful.HashFromPlainText(token)
	t, err := s.repo.GetTokenByHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if t.Scope != string(stateful.PasswordReset) {
		return appErrors.BadData
	}

	if time.Now().After(t.ExpiresAt) || t.RevokedAt != nil {
		return appErrors.Unauthorized
	}

	password := password.Password{}
	err = password.Set(newPassword)
	if err != nil {
		return err
	}

	rowsAffected, err := s.repo.UpdateUserPassword(ctx, repo.UpdateUserPasswordParams{
		ID:           t.UserID,
		PasswordHash: password.Hash,
		Salt:         password.Salt,
	},
	)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return appErrors.NotFound
	}

	return nil
}
