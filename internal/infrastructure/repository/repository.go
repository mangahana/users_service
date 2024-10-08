package repository

import (
	"context"
	"users_service/internal/core"
	"users_service/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepo interface {
	Create(ctx context.Context, phone, username, password string) (int, error)
	CreateSession(ctx context.Context, userId int) (string, error)

	GetIdByCredentials(ctx context.Context, phone, password string) (int, error)
	GetOne(ctx context.Context, userId int) (*models.User, error)
	GetSession(ctx context.Context, token string) (models.Session, error)

	UpdatePhoto(ctx context.Context, userId int, filename string) error
	UpdatePasswordByID(ctx context.Context, userId int, newPassword string) error

	CheckPassword(ctx context.Context, userId int, password string) error

	PhoneExists(ctx context.Context, phone string) (bool, error)
	IsCodeSent(ctx context.Context, phone string) (bool, error)
	CreateConfirmationCode(ctx context.Context, phone, code, ip string) error
	GetSMSbyCredentials(ctx context.Context, dto *core.ConfirmationDTO) error
	DeleteConfirmationCodesByPhone(ctx context.Context, phone string) error
}

type repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) IRepo {
	return &repo{
		db: db,
	}
}
