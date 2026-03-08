package service

import (
	"context"
	"study-stack/internal/adapters/sqlc/repo"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/consts"

	"github.com/google/uuid"
)

func (s *Service) GetAndUseFlashcard(ctx context.Context, userID uuid.UUID) (repo.Flashcard, error) {
	if userID == uuid.Nil {
		return repo.Flashcard{}, appErrors.BadData
	}

	return s.repo.GetAndUseFlashCard(ctx, s.db, userID)
}

type flashcardPage struct {
	Flashcards  []repo.GetFlashcardsPageRow `json:"flashcards"`
	HasNextPage bool                        `json:"has_next_page"`
}
type GetFlashCardPageParams struct {
	UserID            uuid.UUID
	LastSeenFlashcard *uuid.UUID
}

func (s *Service) GetFlashCardPage(ctx context.Context, params GetFlashCardPageParams) (flashcardPage, error) {
	if params.UserID == uuid.Nil {
		return flashcardPage{}, appErrors.BadData
	}

	var cursor uuid.UUID
	if params.LastSeenFlashcard != nil {
		cursor = *params.LastSeenFlashcard
	} else {
		cursor = uuid.Max
	}

	page, err := s.repo.GetFlashcardsPage(ctx, repo.GetFlashcardsPageParams{
		UserID:              params.UserID,
		LastSeenFlashcardID: cursor,
	})
	if err != nil {
		return flashcardPage{}, err
	}

	return flashcardPage{
		Flashcards:  page,
		HasNextPage: len(page) == consts.PAGE_SIZE,
	}, nil
}
