package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) IsCodeSent(ctx context.Context, phone string) (bool, error) {
	var code string
	sql := "SELECT code FROM confirmation_codes WHERE phone = $1 AND created_at + INTERVAL '3 minutes' > NOW();"
	err := r.db.QueryRow(ctx, sql, phone).Scan(&code)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	} else {
		return false, err
	}
}
