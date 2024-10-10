package cerror

func InvalidData() error {
	return &customError{
		Code:    "INVALID_DATA",
		Message: "invalid data",
	}
}
