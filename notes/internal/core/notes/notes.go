package notes

import (
	"context"

	"github.com/google/uuid"
)

// TODO: Note repo should work only with uuidV7
type NoteRepo interface {
	Save(ctx context.Context, note Note) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (*Note, error)
	Update(ctx context.Context, id uuid.UUID) (*Note, error)

	FindConnectedNotes(ctx context.Context, id uuid.UUID)
	FindNotesByTags(ctx context.Context, id uuid.UUID)
}

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
