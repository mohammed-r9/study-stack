package handler

import (
	"log"
	"net/http"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/env"
)

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.URL.Query().Get("t")

	if tokenStr == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println(appErrors.InvalidVerificationToken)
		return
	}

	err := h.svc.VerifyEmail(r.Context(), tokenStr)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, env.Config.FRONTEND_URL, http.StatusFound)

}
