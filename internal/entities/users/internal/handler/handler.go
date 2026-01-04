package handler

import (
	"database/sql"
	"study-stack/internal/entities/users/internal/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		svc: *service.NewService(db),
	}
}
