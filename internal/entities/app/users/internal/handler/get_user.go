package handler

import (
	"encoding/json"
	"net/http"
	"study-stack/internal/shared/utils"
)

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := h.svc.GetUserByID(r.Context(), userData.UserID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
