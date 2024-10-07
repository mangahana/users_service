package repository

import "context"

func (r *repo) GetIdByCredentials(ctx context.Context, phone, password string) (int, error) {
	var userId int
	sql := "SELECT id FROM users WHERE phone = $1 AND password = $2;"
	err := r.db.QueryRow(ctx, sql, phone, password).Scan(&userId)
	return userId, err
}
