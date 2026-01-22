package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/entities/app/collections/service"
	"study-stack/internal/shared/utils"
)

type createReq struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *Handler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	req := createReq{}
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
	err = h.svc.CreateCollection(r.Context(), service.CreateCollectionParams{
		UserID:     userData.UserID,
		Title:      req.Title,
		Desription: req.Description,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Printf("error while creating collection: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
