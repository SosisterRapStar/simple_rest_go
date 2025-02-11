package domain

import (
	"github.com/google/uuid"
)

type Note struct {
	Id      uuid.UUID
	Title   string
	Content string
}

type ErrorCreatingNote struct {
	msg string
}

func (e *ErrorCreatingNote) Error() string {
	if e.msg != "" {
		return e.msg
	}
	return "Erorr occured during: creating note object"
}
