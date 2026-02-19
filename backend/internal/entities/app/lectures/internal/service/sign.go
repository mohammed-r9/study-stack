package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"

	"github.com/google/uuid"
)

type GetSignedURLParams struct {
	UserID    uuid.UUID
	LectureID uuid.UUID
}

func (s *Service) GetSignedURL(ctx context.Context, params GetSignedURLParams) (string, error) {
	lecture, err := s.repo.GetLectureFileKey(ctx, repo.GetLectureFileKeyParams{ID: params.LectureID, UserID: params.UserID})
	if err != nil {
		return "", err
	}
	url, err := s.bucket.GetURL(ctx, lecture.FileKey)
	if err != nil {
		return "", err
	}

	return url, nil
}
