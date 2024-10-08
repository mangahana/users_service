package application

import (
	"context"
	"users_service/internal/core"
	"users_service/internal/infrastructure/repository"
	"users_service/internal/infrastructure/sms"
	"users_service/internal/models"
)

type UseCase interface {
	GetOne(ctx context.Context, userId int) (*models.User, error)

	GetSession(ctx context.Context, token string) (models.Session, error)

	SignIn(ctx context.Context, dto *core.SignInDTO) (string, error)
	Join(ctx context.Context, phone string) error
	Register(ctx context.Context, dto *core.RegisterDTO) (string, error)
	Confirmation(ctx context.Context, dto *core.ConfirmationDTO) error
	PhoneExists(ctx context.Context, phone string) (bool, error)

	UploadPhoto(ctx context.Context, userId int, filename, mime string, data []byte) error
	ChangePassword(ctx context.Context, userId int, dto *core.ChangePasswordDTO) error
}

type useCase struct {
	repo         repository.IRepo
	sms          *sms.SmsService
	uploadFolder string
}

func New(repo repository.IRepo, sms *sms.SmsService, uploadFolder string) *useCase {
	return &useCase{
		repo:         repo,
		sms:          sms,
		uploadFolder: uploadFolder,
	}
}
