package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"study-stack/internal/entities/app/users/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"
)

type UpdateUserRequest struct {
	Name     *string                `json:"name,omitempty"`
	Email    *string                `json:"email,omitempty"`
	Password *UpdatePasswordRequest `json:"password,omitempty"`
}

type UpdatePasswordRequest struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.DataFromContext(r.Context())
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	req := UpdateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("error decoding request: %v\n", err)
		return
	}

	if req.Name == nil && req.Email == nil && req.Password == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if req.Name != nil {
		err := h.validate.Var(*req.Name, "min=2,max=64")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = h.svc.UpdateUserName(r.Context(), service.UpdateNameParams{
			UserID:  userData.UserID,
			NewName: *req.Name,
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	// TODO: send correct status codes, not always 500
	if req.Email != nil {
		err := h.validate.Var(*req.Email, "email")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = h.svc.UpdateUserEmail(r.Context(), service.UpdateEmailParams{
			UserID:   userData.UserID,
			NewEmail: *req.Email,
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	if req.Password != nil {
		err := h.validate.Var(req.Password.New, "min=8")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		err = h.validate.Var(req.Password.Old, "min=8")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		if req.Password.New == req.Password.Old {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println(appErrors.NoChange)
			return
		}

		err = h.svc.UpdateUserPassword(r.Context(), service.UpdatePasswordParams{
			UserID:      userData.UserID,
			NewPassword: req.Password.New,
			OldPassword: req.Password.Old,
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
