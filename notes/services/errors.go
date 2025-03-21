package services

import (
	"errors"
)

var (
	ErrNoteNotFound       = errors.New("Service can not find such note")
	ErrBadRequest         = errors.New("Bad request")
	ErrInternalFailure    = errors.New("Internal erorr")
	ErrTooManyRowsToFetch = errors.New("Too many rows to fetch")
	ErrTimeOutExceeded    = errors.New("Something bad occured and timeout was exceeded")
	ErrInvalidInput       = errors.New("It seems your input values are invalid")
)

type ServiceError struct {
	appError    error
	actualError error
}

func (se ServiceError) GetAppError() error {
	return se.appError
}

func (se ServiceError) GetActualError() error {
	return se.actualError
}

func NewServiceError(appError, actualError error) ServiceError {
	return ServiceError{
		appError:    appError,
		actualError: actualError,
	}
}

func (se ServiceError) Error() string {
	return errors.Join(se.actualError, se.appError).Error()
}
