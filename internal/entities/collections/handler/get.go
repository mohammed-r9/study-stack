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

func (h *Handler) GetCollectionByID(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "bad Request", http.StatusBadRequest)
		return
	}

	collectionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad Request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	collection, err := h.svc.GetCollectionByID(r.Context(), userData.UserID, collectionID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}

}

func (h *Handler) GetCollections(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "bad Request", http.StatusBadRequest)
		return
	}

	archived := r.URL.Query().Get("archived")
	collections := []repo.Collection{}
	var err error

	w.Header().Set("Content-Type", "application/json")
	switch archived {
	case "":
		collections, err = h.svc.GetAllCollections(r.Context(), userData.UserID)
	case "true":
		collections, err = h.svc.GetAllArchived(r.Context(), userData.UserID)
	case "false":
		collections, err = h.svc.GetAllUnarchived(r.Context(), userData.UserID)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("invalid filter in collections")
		return
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(collections); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}
