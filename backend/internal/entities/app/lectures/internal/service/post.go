package service

import (
	"context"
	"mime/multipart"
	"strings"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (s *Service) InsertLecture(ctx context.Context, userID, materialID uuid.UUID, lectureTitle string, file *multipart.FileHeader) error {
	if userID == uuid.Nil || materialID == uuid.Nil || lectureTitle == "" {
		return appErrors.BadData
	}
	if file.Size > 20*1024*1024 {
		return appErrors.FileTooLarge
	}
	if !strings.HasSuffix(file.Filename, ".pdf") {
		return appErrors.BadData
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileKey, err := uuid.NewV7()
	if err != nil {
		return err
	}
	err = s.bucket.Upload(ctx, fileKey.String(), src, "application/pdf")
	if err != nil {
		return err
	}

	lectureID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	err = s.repo.CreateLecture(ctx, repo.CreateLectureParams{
		UserID:     userID,
		ID:         lectureID,
		Title:      lectureTitle,
		MaterialID: materialID,
		FileKey:    fileKey,
		FileSize:   file.Size,
	})

	return err
}
