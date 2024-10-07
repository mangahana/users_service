package application

import "context"

func (u *useCase) PhoneExists(ctx context.Context, phone string) (bool, error) {
	return u.repo.PhoneExists(ctx, phone)
}
