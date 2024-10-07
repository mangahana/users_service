package application

import (
	"context"
	"users_service/internal/core"
	"users_service/internal/infrastructure/repository"
	"users_service/internal/infrastructure/sms"
)

type UseCase interface {
	// SignIn(ctx context.Context, dto *core.SignInDTO) (string, error)
	Join(ctx context.Context, phone string) error
	Register(ctx context.Context, dto *core.RegisterDTO) (string, error)
	Confirmation(ctx context.Context, dto *core.ConfirmationDTO) error
	PhoneExists(ctx context.Context, phone string) (bool, error)
}

type useCase struct {
	repo repository.IRepo
	sms  *sms.SmsService
}

func New(repo repository.IRepo, sms *sms.SmsService) *useCase {
	return &useCase{
		repo: repo,
		sms:  sms,
	}
}
