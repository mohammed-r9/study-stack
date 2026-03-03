package service

import (
	"database/sql"
	"study-stack/internal/adapters/sqlc/repo"
)

type Service struct {
	repo repo.Queries
	db   *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: *repo.New(db),
		db:   db,
	}
}
