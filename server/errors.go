package server

import "errors"

var (
	BadRequest   = errors.New("Bad request")
	Unauthorized = errors.New("Unauthorized")
	Forbidden    = errors.New("Forbidden request")
	NotFound     = errors.New("Not found")
)

type Error struct {
	appError error
	svcError error
}

func (e Error) Error() string {
	return errors.Join(e.svcError, e.appError).Error()
}

func NewError(svcErr, appErr error) error {
	return Error{
		svcError: svcErr,
		appError: appErr,
	}
}
