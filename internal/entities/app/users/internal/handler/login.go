package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/entities/app/users/internal/service"
	"study-stack/internal/shared/utils"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"` // maybe make this validation better? idk, it's a toy project after all
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	req := loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}
	err := h.validate.Struct(req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("invalid request: %v\n", err)
		return
	}

	tokens, err := h.svc.Login(r.Context(), service.LoginParams{
		Email:       req.Email,
		Password:    req.Password,
		Device_name: utils.GetDeviceNameFromUserAgent(r.UserAgent()),
	})

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error logging in: %v\n", err)
		return
	}

	utils.SetRefreshCookie(w, tokens.Refresh)
	utils.SetCsrfCookie(w, tokens.Csrf)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": tokens.Access,
	})
}
