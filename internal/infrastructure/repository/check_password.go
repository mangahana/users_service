package repository

import "context"

func (r *repo) CheckPassword(ctx context.Context, userId int, password string) error {
	var id int
	sql := "SELECT id FROM users WHERE id = $1 AND password = $2"
	return r.db.QueryRow(ctx, sql, userId, password).Scan(&id)
}
