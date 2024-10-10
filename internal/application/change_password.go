package application

import (
	"context"
	"users_service/internal/core"
	"users_service/internal/core/cerror"
)

func (u *useCase) ChangePassword(ctx context.Context, userId int, dto *core.ChangePasswordDTO) error {
	if err := u.repo.CheckPassword(ctx, userId, dto.OldPassword); err != nil {
		return cerror.New(cerror.INVALID_PASSWORD, "old password is invalid")
	}

	return u.repo.UpdatePasswordByID(ctx, userId, dto.NewPassword)
}
