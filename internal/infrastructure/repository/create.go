package repository

import "context"

func (r *repo) Create(ctx context.Context, phone, username, password string) (int, error) {
	var userId int
	sql := "INSERT INTO users (phone, username, password) VALUES($1, $2, $3) RETURNING id;"
	err := r.db.QueryRow(ctx, sql, phone, username, password).Scan(&userId)
	return userId, err
}
