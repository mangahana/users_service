package application

import (
	"context"
	"users_service/internal/models"
)

func (u *useCase) GetOne(ctx context.Context, userId int) (*models.User, error) {
	return u.repo.GetOne(ctx, userId)
}
