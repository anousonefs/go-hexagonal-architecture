package errs

import "net/http"

type AppError struct {
	Code int
	Message string
}

func (e AppError) Error() string{
	return e.Message
}

func NewNotFoundError(message string) error{
	return AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error{
	return AppError{
		Code: http.StatusInternalServerError,
		Message: "Unexpected Error!!",
	}
}

func NewValidateError(message string) error{
	return AppError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}