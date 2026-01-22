package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/adapters/sqlc/repo"
	"study-stack/internal/shared/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type getAllReq struct {
	CollectionID uuid.UUID `json:"collection_id" validate:"required"`
}

func (h *Handler) GetAllMaterials(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	req := getAllReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}
	err := h.validate.Struct(req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	filter := r.URL.Query().Get("archived")

	materials := []repo.Material{}
	switch filter {
	case "true":
		materials, err = h.svc.GetAllArchived(r.Context(), userData.UserID, req.CollectionID)
	case "false":
		materials, err = h.svc.GetAllUnarchived(r.Context(), userData.UserID, req.CollectionID)
	case "":
		materials, err = h.svc.GetAll(r.Context(), userData.UserID, req.CollectionID)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("invalid filter in materials")
		return
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(materials); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func (h *Handler) GetMaterialByID(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	idStr := chi.URLParam(r, "id")
	materialID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	material, err := h.svc.GetByID(r.Context(), userData.UserID, materialID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(material); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}
