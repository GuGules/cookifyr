package apperror

import "net/http"

type AppError struct {
	Status int
	Msg    string
}

func (appErr AppError) Error() string {
	return appErr.Msg
}

func NewInternalServerError() AppError {
	return AppError{
		Status: http.StatusInternalServerError,
		Msg:    "internal server error",
	}
}
