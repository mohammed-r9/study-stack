package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

// lastSeenID is nil when it's the first page
func (s *Service) GetAllLectures(ctx context.Context, userID uuid.UUID, lastSeenID *uuid.UUID) ([]repo.Lecture, error) {
	if userID == uuid.Nil {
		return []repo.Lecture{}, appErrors.BadData
	}

	lastLectureID := uuid.Nil
	if lastSeenID != nil {
		lastLectureID = *lastSeenID
	}

	return s.repo.ListLectures(ctx, repo.ListLecturesParams{UserID: userID, LastSeenLectureID: lastLectureID})
}
