package repository

import (
	"context"
	"log"
	"users_service/internal/core"

	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateSession(ctx context.Context, userId int) (string, error) {
	token, err := core.GenerateRandomString(64)
	if err != nil {
		return "", err
	}

	sql := "INSERT INTO sessions (user_id, token) VALUES($1, $2);"
	err = r.db.QueryRow(ctx, sql, userId, token).Scan()
	if err != nil && err != pgx.ErrNoRows {
		log.Println("here")
		return "", err
	}

	return token, nil
}
