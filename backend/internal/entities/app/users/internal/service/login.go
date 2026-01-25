package service

import (
	"context"
	"database/sql"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/entities/tokens/stateful"
	"study-stack/internal/entities/tokens/stateless"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/password"

	"github.com/google/uuid"
)

type LoginParams struct {
	Email       string
	Password    string
	Device_name string
}

type loginTokens struct {
	Access  string
	Refresh string
	Csrf    string
}

func (s *Service) Login(ctx context.Context, params LoginParams) (loginTokens, error) {
	user, err := s.repo.GetUserByEmail(ctx, params.Email)
	if err == sql.ErrNoRows {
		return loginTokens{}, appErrors.InvalidCredentials
	}
	if err != nil {
		return loginTokens{}, err
	}

	password := password.Password{
		Hash: user.PasswordHash,
		Salt: user.Salt,
	}

	err = password.Matches(params.Password)
	if err != nil {
		return loginTokens{}, appErrors.InvalidCredentials
	}

	refreshToken, err := stateful.NewRefreshToken()
	if err != nil {
		return loginTokens{}, err
	}
	isVerified := false
	if user.VerifiedAt != nil {
		isVerified = true
	}

	accessToken, err := stateless.NewAcessToken(stateless.UserClaims{UserID: user.ID, IsVerified: isVerified})
	if err != nil {
		return loginTokens{}, err
	}

	err = s.repo.NewUserSession(ctx, repo.NewUserSessionParams{
		ID:         uuid.New(),
		UserID:     user.ID,
		TokenHash:  refreshToken.Hash,
		CsrfHash:   refreshToken.CsrfHash,
		DeviceName: params.Device_name,
	})
	if err != nil {
		return loginTokens{}, err
	}

	return loginTokens{
		Refresh: refreshToken.PlainText,
		Csrf:    refreshToken.CsrfPlainText,
		Access:  accessToken,
	}, nil

}
