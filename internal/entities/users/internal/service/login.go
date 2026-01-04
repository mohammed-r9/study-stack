package service

import (
	"context"
	"study-stack/internal/entities/tokens/stateful"
	"study-stack/internal/entities/tokens/stateless"
	"study-stack/internal/shared/password"
)

type LoginParams struct {
	Email    string
	Password string
}

type loginTokens struct {
	Access  string
	Refresh string
	Csrf    string
}

func (s *Service) Login(ctx context.Context, params LoginParams) (loginTokens, error) {
	user, err := s.repo.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return loginTokens{}, err
	}

	password := password.Password{
		Hash: user.PasswordHash,
		Salt: user.Salt,
	}

	err = password.Matches(params.Password)
	if err != nil {
		return loginTokens{}, err
	}

	refreshToken, err := stateful.NewRefreshToken()
	if err != nil {
		return loginTokens{}, err
	}
	isVerified := false
	if user.VerifiedAt != nil {
		isVerified = false
	}
	accessToken, err := stateless.NewAcessToken(stateless.UserClaims{UserID: user.ID, IsVerified: isVerified})
	if err != nil {
		return loginTokens{}, err
	}
	return loginTokens{
		Refresh: refreshToken.PlainText,
		Csrf:    refreshToken.CsrfPlainText,
		Access:  accessToken,
	}, nil

}
