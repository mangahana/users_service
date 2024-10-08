package repository

import (
	"context"
	"users_service/internal/models"
)

func (r *repo) GetOne(ctx context.Context, userId int) (*models.User, error) {
	var output models.User
	sql := "SELECT id, username, description, photo FROM users WHERE id = $1;"
	err := r.db.QueryRow(ctx, sql, userId).Scan(&output.ID, &output.Username, &output.Description, &output.Photo)
	return &output, err
}
