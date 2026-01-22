package service

import (
	"context"
	"study-stack/internal/entities/tokens/stateful"
	"study-stack/internal/entities/tokens/stateless"
	appErrors "study-stack/internal/shared/app_errors"
)

func (s *Service) RefreshToken(ctx context.Context, refreshPlain, csrfPlain string) (string, error) {
	refreshHash := stateful.HashFromPlainText(refreshPlain)
	session, err := s.repo.GetSessionByHash(ctx, refreshHash)
	if err != nil {
		return "", err
	}

	if !stateful.CompareOpaqueTokens(refreshPlain, session.TokenHash) {
		return "", appErrors.InvalidRefreshToken
	}

	if !stateful.CompareOpaqueTokens(csrfPlain, session.CsrfHash) {
		return "", appErrors.InvalidCsrfToken
	}

	jwt, err := stateless.NewAcessToken(stateless.UserClaims{
		UserID:     session.UserID,
		IsVerified: session.VerifiedAt != nil,
	})

	if err != nil {
		return "", err
	}

	return jwt, nil
}
