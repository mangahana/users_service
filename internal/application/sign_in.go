package application

import (
	"context"
	"users_service/internal/core"
)

func (u *useCase) SignIn(ctx context.Context, dto *core.SignInDTO) (string, error) {
	userId, err := u.repo.GetIdByCredentials(ctx, dto.Phone, dto.Password)
	if err != nil {
		return "", err
	}

	return u.repo.CreateSession(ctx, userId)
}
