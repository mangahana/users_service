package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateConfirmationCode(ctx context.Context, phone, code, ip string) error {
	sql := "INSERT INTO confirmation_codes (code, phone, ip) VALUES($1, $2, $3);"
	err := r.db.QueryRow(ctx, sql, code, phone, ip).Scan()

	if errors.Is(err, pgx.ErrNoRows) || err == nil {
		return nil
	}

	return err
}
