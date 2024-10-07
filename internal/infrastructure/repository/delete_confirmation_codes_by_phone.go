package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) DeleteConfirmationCodesByPhone(ctx context.Context, phone string) error {
	sql := "DELETE FROM confirmation_codes WHERE phone = $1"
	err := r.db.QueryRow(ctx, sql, phone).Scan()
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}

	return err
}
