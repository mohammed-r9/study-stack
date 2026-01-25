package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (q *Queries) VerifyUser(ctx context.Context, db *sql.DB, userID uuid.UUID, tokenHash string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := q.WithTx(tx)

	err = qtx.verifyUserEmail(ctx, userID)
	if err != nil {
		return err
	}

	_, err = qtx.UseToken(ctx, tokenHash)
	if err != nil {
		return err
	}

	return tx.Commit()
}
