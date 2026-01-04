package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/entities/users/internal/service"

	"github.com/go-playground/validator/v10"
)

type registerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"` // maybe make this validation better? idk, it's a toy project after all
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	req := registerRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}
	err := validate.Struct(req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("invalid request: %v\n", err)
		return
	}

	err = h.svc.RegisterUser(r.Context(), service.RegisterParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error registering user: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
