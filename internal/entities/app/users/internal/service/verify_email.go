package service

import (
	"context"
	"study-stack/internal/entities/tokens/stateful"
	appErrors "study-stack/internal/shared/app_errors"
	"time"
)

func (s *Service) VerifyEmail(ctx context.Context, token string) error {
	if token == "" {
		return appErrors.BadRequest
	}

	tokenHash := stateful.HashFromPlainText(token)
	storedToken, err := s.repo.GetTokenByHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if time.Now().After(storedToken.ExpiresAt) || storedToken.UsedAt != nil {
		return appErrors.Unauthorized
	}

	err = s.repo.VerifyUser(ctx, s.db, storedToken.UserID, storedToken.Hash)
	if err != nil {
		return err
	}

	return nil
}
