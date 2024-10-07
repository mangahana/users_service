package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) PhoneExists(ctx context.Context, phone string) (bool, error) {
	var userId int

	sql := "SELECT id FROM users WHERE phone = $1"
	err := r.db.QueryRow(ctx, sql, phone).Scan(&userId)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	} else {
		return false, err
	}
}
