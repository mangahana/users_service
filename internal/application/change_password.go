package application

import (
	"context"
	"users_service/internal/core"
)

func (u *useCase) ChangePassword(ctx context.Context, userId int, dto *core.ChangePasswordDTO) error {
	if err := u.repo.CheckPassword(ctx, userId, dto.OldPassword); err != nil {
		return err
	}

	return u.repo.UpdatePasswordByID(ctx, userId, dto.NewPassword)
}
