package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Println(err.Error())
		return
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("no csrf present")
		return
	}

	accessToken, err := h.svc.RefreshToken(r.Context(), refreshCookie.Value, csrf)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error generating the jwt: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})

}
