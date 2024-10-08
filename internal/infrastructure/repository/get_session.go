package repository

import (
	"context"
	"users_service/internal/models"
)

func (r *repo) GetSession(ctx context.Context, token string) (models.Session, error) {
	var output models.Session
	sql := `SELECT id, is_banned,
	(SELECT array_agg(permissions.name) FROM roles
		LEFT JOIN permissions ON permissions.id = any(roles.permissions)
	WHERE roles.id = role_id) as permissions
	FROM users WHERE id = (SELECT user_id FROM sessions WHERE token = $1)`

	err := r.db.QueryRow(ctx, sql, token).Scan(&output.UserID, &output.IsBanned, &output.Permissions)
	return output, err
}
