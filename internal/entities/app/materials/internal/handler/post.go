package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/shared/utils"

	"github.com/google/uuid"
)

type materialCreationReq struct {
	CollectionID uuid.UUID `json:"collection_id" validate:"required"`
	Title        string    `json:"title" validate:"required,min=4"`
}

func (h *Handler) InsertMaterial(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	req := materialCreationReq{}
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

	err = h.svc.InsertMaterial(r.Context(), req.Title, userData.UserID, req.CollectionID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
