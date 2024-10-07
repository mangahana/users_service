package application

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"users_service/internal/core"
)

func (u *useCase) Join(ctx context.Context, phone string) error {
	phoneExists, err := u.repo.PhoneExists(ctx, phone)
	if err != nil {
		return err
	}
	if phoneExists {
		return core.ErrPhoneIsAlreadyInUse
	}

	isSent, err := u.repo.IsCodeSent(ctx, phone)
	if err != nil {
		log.Println(err)
		return err
	}
	if isSent {
		return core.ErrCodeAlreadySent
	}

	confirmationCode, err := Code(6)
	if err != nil {
		return err
	}

	err = u.sms.Send("7"+phone, "semser.org\nРастау коды: "+confirmationCode)
	if err != nil {
		return err
	}

	err = u.repo.CreateConfirmationCode(ctx, phone, confirmationCode, "127.0.0.1")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
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
