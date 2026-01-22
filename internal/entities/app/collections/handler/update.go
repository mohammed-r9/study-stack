package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/shared/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type updateReq struct {
	ToArchive   *bool   `json:"to_archive"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (h *Handler) UpdateCollection(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	req := updateReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}

	collectionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error parsing collection id: %v\n", err)
		return
	}

	if req.Title != nil {
		err = h.svc.UpdateTitle(r.Context(), collectionID, userData.UserID, *req.Title)
		return
	}

	if req.Description != nil {
		err = h.svc.UpdateDescription(r.Context(), collectionID, userData.UserID, *req.Description)
		return
	}

	if req.ToArchive != nil {
		err = h.svc.UpdateIsArchived(r.Context(), collectionID, userData.UserID, *req.ToArchive)
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("error updating collection: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
