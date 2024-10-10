package application

import (
	"context"
	"crypto/rand"
	"math/big"
	"users_service/internal/core/cerror"
)

func (u *useCase) Join(ctx context.Context, phone string) error {
	phoneExists, err := u.repo.PhoneExists(ctx, phone)
	if err != nil {
		return err
	}
	if phoneExists {
		return cerror.New(cerror.PHONE_USED, "the phone is already in use")
	}

	isSent, err := u.repo.IsCodeSent(ctx, phone)
	if err != nil {
		return err
	}
	if isSent {
		return cerror.New(cerror.CODE_SENT, "code has been sent")
	}

	confirmationCode, err := Code(6)
	if err != nil {
		return err
	}

	err = u.sms.Send("7"+phone, "semser.org\nРастау коды: "+confirmationCode)
	if err != nil {
		return err
	}

	return u.repo.CreateConfirmationCode(ctx, phone, confirmationCode, "127.0.0.1")
}

func Code(length int) (string, error) {
	var output string

	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return "", err
		}
		output += nBig.String()
	}

	return output, nil
}
