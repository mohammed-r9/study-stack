package handler

import (
	"database/sql"
	"study-stack/internal/entities/app/lectures/internal/service"
	S3 "study-stack/internal/s3"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc      service.Service
	validate *validator.Validate
}

func NewHandler(db *sql.DB, validator *validator.Validate, b S3.Storage) *Handler {
	return &Handler{
		svc:      *service.NewService(db, b),
		validate: validator,
	}
}
