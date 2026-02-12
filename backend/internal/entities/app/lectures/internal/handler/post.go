package handler

import (
	"log"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) InsertLecture(c *fiber.Ctx) error {
	userData, ok := utils.DataFromLocals(c)
	if !ok {
		log.Println("invalid user")
		return appErrors.BadData
	}

	materialIDStr := c.FormValue("material_id")
	materialID, err := uuid.Parse(materialIDStr)
	if err != nil {
		log.Println(err)
		return appErrors.BadData
	}
	lectureTitle := c.FormValue("lecture_title")
	if lectureTitle == "" {
		log.Println("invalid title")
		return appErrors.BadData
	}

	file, err := c.FormFile("lecture_file")
	if err != nil {
		log.Println("invalid file")
		return appErrors.BadData
	}
	err = h.svc.InsertLecture(c.Context(), userData.UserID, materialID, lectureTitle, file)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.Status(fiber.StatusCreated).SendString("Lecture uploaded successfully")
}
