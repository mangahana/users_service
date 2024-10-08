package repository

import "context"

func (r *repo) UpdatePasswordByID(ctx context.Context, userId int, newPassword string) error {
	var id int
	sql := "UPDATE users SET password = $1 WHERE id = $2 RETURNING id;"
	return r.db.QueryRow(ctx, sql, newPassword, userId).Scan(&id)
}
