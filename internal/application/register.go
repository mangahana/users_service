package application

import (
	"context"
	"log"
	"regexp"
	"users_service/internal/core"
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
		log.Println("hereeeeeeeee")
		return "", core.ErrInvalidUsername
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
