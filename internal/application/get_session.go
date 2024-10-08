package application

import (
	"context"
	"users_service/internal/models"
)

func (u *useCase) GetSession(ctx context.Context, token string) (models.Session, error) {
	return u.repo.GetSession(ctx, token)
}
