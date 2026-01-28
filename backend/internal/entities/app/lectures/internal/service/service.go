package service

import (
	"database/sql"
	"study-stack/internal/adapters/sqlc/repo"
	S3 "study-stack/internal/s3"
)

type Service struct {
	repo   repo.Queries
	db     *sql.DB
	bucket S3.Storage
}

func NewService(db *sql.DB, bucket S3.Storage) *Service {
	return &Service{
		repo:   *repo.New(db),
		db:     db,
		bucket: bucket,
	}
}
