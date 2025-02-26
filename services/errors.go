package services

import (
	"errors"
)

var (
	ErrNoteNotFound       = errors.New("Service did not find such note")
	ErrBadRequest         = errors.New("Bad request")
	ErrInternalFailure    = errors.New("Internal erorr")
	ErrTooManyRowsToFetch = errors.New("Too many rows to fetch")
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
