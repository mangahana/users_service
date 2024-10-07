package core

type SignInDTO struct {
	Phone    string `json:"phone" validate:"required,min=10,max=10"`
	Password string `json:"password"`
}

type JoinDTO struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
}

type RegisterDTO struct {
	Phone            string `json:"phone" validate:"required,min=10,max=10"`
	ConfirmationCode string `json:"confirmation_code" validate:"required,min=6,max=6"`
	Username         string `json:"username" validate:"required"`
	Password         string `json:"password" validate:"required,min=8"`
}

type ConfirmationDTO struct {
	Phone            string `json:"phone" validate:"required,min=10,max=10"`
	ConfirmationCode string `json:"confirmation_code" validate:"required,min=6,max=6"`
}
