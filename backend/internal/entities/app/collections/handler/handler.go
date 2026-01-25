package handler

import (
	"database/sql"
	"study-stack/internal/entities/app/collections/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc      service.Service
	validate *validator.Validate
}

func NewHandler(db *sql.DB, validator *validator.Validate) *Handler {
	return &Handler{
		svc:      *service.NewService(db),
		validate: validator,
	}
}
