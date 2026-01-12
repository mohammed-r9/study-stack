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
	ToArchive *bool   `json:"to_archive"`
	Title     *string `json:"title"`
}

func (h *Handler) UpdateMaterial(w http.ResponseWriter, r *http.Request) {
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

	materialID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error parsing material id: %v\n", err)
		return
	}

	if req.Title != nil {
		err = h.svc.UpdateMaterialTitle(r.Context(), *req.Title, materialID, userData.UserID)
	}

	if req.ToArchive != nil {
		err = h.svc.UpdateMaterialArchivedAt(r.Context(), *req.ToArchive, materialID, userData.UserID)
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error updating material: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
