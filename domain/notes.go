package domain

import (
	"context"

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

type NoteService interface {
	CreateNote(ctx context.Context, note *Note) (uuid.UUID, error)
	FindNote(ctx context.Context, id string) (*Note, error)
	DeleteNote(ctx context.Context, id string) error
	UpdateNote(ctx context.Context, upd Note) (*Note, error)
}
