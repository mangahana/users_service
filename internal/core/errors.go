package core

import "errors"

var (
	ErrPhoneIsAlreadyInUse = errors.New("this phone is already in use")
	ErrCodeAlreadySent     = errors.New("code for this phone number already been sent")
	ErrInvalidUsername     = errors.New("invalid username format")
)
