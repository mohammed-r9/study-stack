package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/password"

	"github.com/google/uuid"
)

type UpdateNameParams struct {
	UserID  uuid.UUID
	NewName string
}

type UpdateEmailParams struct {
	UserID   uuid.UUID
	NewEmail string
}

type UpdatePasswordParams struct {
	UserID      uuid.UUID
	OldPassword string
	NewPassword string
}

func (s *Service) UpdateUserName(ctx context.Context, params UpdateNameParams) error {
	return s.repo.UpdateUserName(ctx, repo.UpdateUserNameParams{
		ID:   params.UserID,
		Name: params.NewName,
	})

}

func (s *Service) UpdateUserEmail(ctx context.Context, params UpdateEmailParams) error {
	return s.repo.UpdateUserEmail(ctx, repo.UpdateUserEmailParams{
		ID:    params.UserID,
		Email: params.NewEmail,
	})
}

func (s *Service) UpdateUserPassword(ctx context.Context, params UpdatePasswordParams) error {
	user, err := s.repo.GetUserByID(ctx, params.UserID)
	if err != nil {
		return err
	}
	oldPassword := password.Password{
		Hash: user.PasswordHash,
		Salt: user.Salt,
	}
	err = oldPassword.Matches(params.OldPassword)
	if err != nil {
		return appErrors.Unauthorized
	}

	newPassword := password.Password{}
	err = newPassword.Set(params.NewPassword)
	if err != nil {
		return err
	}

	err = s.repo.UpdateUserPassword(ctx, repo.UpdateUserPasswordParams{
		PasswordHash: newPassword.Hash,
		Salt:         newPassword.Salt,
		ID:           params.UserID,
	})

	return err
}
