package services

import (
	"errors"
)

var (
	ErrNoteNotFound    = errors.New("Service did not find such note")
	ErrBadRequest      = errors.New("Bad request")
	ErrInternalFailure = errors.New("Internal erorr")
)

// set fo erros for driving sides (higher level)
type ServiceError struct {
	// this error is app specific and valid for any transport, only transport (driving side) know how to handle it and what error code to use
	appError error
	// this is low level error, implementation specific error i.g. (MongoDB error, Redis Cache error, Postgres tx error)
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
