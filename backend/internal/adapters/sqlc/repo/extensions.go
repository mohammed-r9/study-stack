package repo

import (
	"context"
	"database/sql"
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/google/uuid"
)

func (q *Queries) VerifyUser(ctx context.Context, db *sql.DB, userID uuid.UUID, tokenHash string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := q.WithTx(tx)

	rowsAffected, err := qtx.verifyUserEmail(ctx, userID)

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return appErrors.NotFound
	}

	_, err = qtx.UseToken(ctx, tokenHash)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// returns the oldest flashcard and updates its last_used field to the current time
func (q *Queries) GetAndUseFlashCard(ctx context.Context, db *sql.DB, userID uuid.UUID) (Flashcard, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return Flashcard{}, err
	}
	defer tx.Rollback()

	qtx := q.WithTx(tx)

	flashcard, err := qtx.getOldestFlashcard(ctx, userID)
	if err != nil {
		return Flashcard{}, err
	}

	rowsAfected, err := qtx.useFlashcard(ctx, flashcard.ID)
	if err != nil {
		return Flashcard{}, err
	}

	if rowsAfected == 0 {
		return Flashcard{}, appErrors.NoRowsAffected
	}

	return flashcard, tx.Commit()
}

type FlashcardImageParams struct {
	Flashcard createImageFlashcardParams
	Labels    []insertImageLabelParams
}

type insertedImageFlashcard struct {
	Image  ImageFlashcard
	Labels []ImageLabel
}

func (q *Queries) InsertImageFlashcardAndLabels(ctx context.Context, db *sql.DB, params FlashcardImageParams) (insertedImageFlashcard, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return insertedImageFlashcard{}, err
	}
	defer tx.Rollback()

	qtx := q.WithTx(tx)

	flashcard, err := qtx.createImageFlashcard(ctx, params.Flashcard)
	if err != nil {
		return insertedImageFlashcard{}, err
	}
	card := insertedImageFlashcard{}
	card.Image = flashcard
	for _, label := range params.Labels {
		insertedLabel, err := qtx.insertImageLabel(ctx, label)
		if err != nil {
			return insertedImageFlashcard{}, err
		}
		card.Labels = append(card.Labels, insertedLabel)
	}

	return card, tx.Commit()
}
