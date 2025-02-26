package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Note struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type UpdateNote struct {
	Title   *string
	Content *string
}

func (n *Note) String() string {
	return fmt.Sprintf("Note{Id: %s, Title: %s, Content: %s}", n.Id, n.Title, n.Content)
}

func NewNote(title string, content string) *Note {
	return &Note{
		Title:   title,
		Content: content}
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
	CreateNote(ctx context.Context, note *Note) (string, error)
	GetNote(ctx context.Context, id string) (*Note, error)
	DeleteNote(ctx context.Context, id string) (string, error)
	UpdateNote(ctx context.Context, upd *UpdateNote, id string) (*Note, error)
	FindNotes(ctx context.Context) ([]Note, error)
}
