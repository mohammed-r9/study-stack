package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Email string `json:"email"`
}

func (h *Handler) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	req := request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}
	token, err := h.svc.RequestPasswordReset(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("error generating password reset token: %v\n", err)
		return
	}
	_ = token
	// need to send the token on the email later

	w.WriteHeader(http.StatusOK)

}
