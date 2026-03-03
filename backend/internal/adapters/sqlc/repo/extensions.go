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

// return the oldest flashcard and updates its last_used field to the current time
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
