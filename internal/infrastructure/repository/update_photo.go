package repository

import "context"

func (r *repo) UpdatePhoto(ctx context.Context, userId int, filename string) error {
	var id int
	sql := "UPDATE users SET photo = $1 WHERE id = $2 RETURNING id;"
	err := r.db.QueryRow(ctx, sql, filename, userId).Scan(&id)
	return err
}
