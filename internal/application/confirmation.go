package application

import (
	"context"
	"users_service/internal/core"
)

func (u *useCase) Confirmation(ctx context.Context, dto *core.ConfirmationDTO) error {
	return u.repo.GetSMSbyCredentials(ctx, dto)
}
