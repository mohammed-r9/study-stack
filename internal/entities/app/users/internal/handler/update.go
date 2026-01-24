package handler

import (
	"log"
	"study-stack/internal/entities/app/users/internal/service"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
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

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		return appErrors.BadRequest
	}

	req := new(UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		log.Printf("error decoding request: %v\n", err)
		return appErrors.BadRequest
	}

	if req.Name == nil && req.Email == nil && req.Password == nil {
		return appErrors.BadRequest
	}

	// update name
	if req.Name != nil {
		if err := h.validate.Var(*req.Name, "min=2,max=64"); err != nil {
			log.Println(err)
			return appErrors.BadRequest
		}
		if err := h.svc.UpdateUserName(c.Context(), service.UpdateNameParams{
			UserID:  userData.UserID,
			NewName: *req.Name,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	// update email
	if req.Email != nil {
		if err := h.validate.Var(*req.Email, "email"); err != nil {
			log.Println(err)
			return appErrors.BadRequest
		}
		if err := h.svc.UpdateUserEmail(c.Context(), service.UpdateEmailParams{
			UserID:   userData.UserID,
			NewEmail: *req.Email,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	// update password
	if req.Password != nil {
		if err := h.validate.Var(req.Password.New, "min=8"); err != nil {
			log.Println(err)
			return appErrors.BadRequest
		}
		if err := h.validate.Var(req.Password.Old, "min=8"); err != nil {
			log.Println(err)
			return appErrors.BadRequest
		}
		if req.Password.New == req.Password.Old {
			log.Println(appErrors.NoChange)
			return appErrors.BadRequest
		}
		if err := h.svc.UpdateUserPassword(c.Context(), service.UpdatePasswordParams{
			UserID:      userData.UserID,
			NewPassword: req.Password.New,
			OldPassword: req.Password.Old,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
