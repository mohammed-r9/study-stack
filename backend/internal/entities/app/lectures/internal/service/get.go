package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/consts"

	"github.com/google/uuid"
)

type lecturesPage struct {
	Lectures    []repo.ListLecturesRow `json:"lectures"`
	HasNextPage bool                   `json:"has_next_page"`
}

type GetAllLecturesParams struct {
	UserID     uuid.UUID
	MaterialID uuid.UUID
	LastSeenID *uuid.UUID
}

// lastSeenID is nil when it's the first page
func (s *Service) GetAllLectures(ctx context.Context, params GetAllLecturesParams) (lecturesPage, error) {
	if params.UserID == uuid.Nil || params.MaterialID == uuid.Nil {
		return lecturesPage{}, appErrors.BadData
	}

	var cursor uuid.UUID
	if params.LastSeenID != nil {
		cursor = *params.LastSeenID
	} else {
		cursor = uuid.Max
	}

	lectures, err := s.repo.ListLectures(ctx, repo.ListLecturesParams{
		UserID:            params.UserID,
		MaterialID:        params.MaterialID,
		LastSeenLectureID: cursor,
	})
	if err != nil {
		return lecturesPage{}, err
	}

	return lecturesPage{
		Lectures:    lectures,
		HasNextPage: len(lectures) == consts.PAGE_SIZE,
	}, nil
}
