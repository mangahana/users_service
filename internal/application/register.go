package application

import (
	"context"
	"regexp"
	"users_service/internal/core"
	"users_service/internal/core/cerror"
)

func (u *useCase) Register(ctx context.Context, dto *core.RegisterDTO) (string, error) {
	confirmationDto := &core.ConfirmationDTO{
		ConfirmationCode: dto.ConfirmationCode,
		Phone:            dto.Phone,
	}
	if err := u.Confirmation(ctx, confirmationDto); err != nil {
		return "", err
	}

	regex, err := regexp.Compile("^[a-zA-Z0-9]+(_?[a-zA-Z0-9]+)*$")
	if err != nil {
		return "", err
	}

	if !regex.Match([]byte(dto.Username)) {
		return "", cerror.New(cerror.USERNAME_INVALID, "username is not valid")
	}

	if len(dto.Password) < 8 {
		return "", cerror.New(cerror.PASSWORD_TOO_SHORT, "the password is too short")
	}

	userId, err := u.repo.Create(ctx, dto.Phone, dto.Username, dto.Password)
	if err != nil {
		return "", err
	}

	if err := u.repo.DeleteConfirmationCodesByPhone(ctx, dto.Phone); err != nil {
		return "", err
	}

	return u.repo.CreateSession(ctx, userId)
}
