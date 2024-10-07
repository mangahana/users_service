package repository

import (
	"context"
	"users_service/internal/core"
)

func (r *repo) GetSMSbyCredentials(ctx context.Context, dto *core.ConfirmationDTO) error {
	var code string
	sql := "SELECT code FROM confirmation_codes WHERE code = $1 AND phone = $2 AND created_at + INTERVAL '3 minutes' > NOW();"
	return r.db.QueryRow(ctx, sql, dto.ConfirmationCode, dto.Phone).Scan(&code)
}
