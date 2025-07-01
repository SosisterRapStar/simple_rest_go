package notes

import (
	"context"

	"github.com/google/uuid"
	"github.com/sosisterrapstar/simple_rest_go/internal/core"
)

// TODO: Note repo should work only with uuidV7
type NoteRepo interface {
	Save(ctx context.Context, params *CreateNote) (*Note, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (*Note, error)
	Update(ctx context.Context, params *UpdateNote) (*Note, error)
	List(ctx context.Context, id uuid.UUID, paging core.LimitOffsetPaging) ([]*Note, error)

	// FindConnectedNotes(ctx context.Context, id uuid.UUID) ([]*Note, error)
	// FindNotesByTags(ctx context.Context, id uuid.UUID, tags []Tag) ([]*Note, error)
}

type Module struct {
	repo NoteRepo
}

func (n *Module) CreateNewNote(ctx context.Context, params *CreateNote) (*Note, error) {
	note, err := n.repo.Save(ctx, params)
	if err != nil {
		return nil, nil
	}
	return note, nil
}

func (n *Module) DeleteNote(ctx context.Context, id uuid.UUID) error {
	err := n.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (n *Module) UpdateNote(ctx context.Context, params *UpdateNote) (*Note, error) {
	note, err := n.repo.Update(ctx, params)
	if err != nil {
		return nil, err
	}
	return note, nil
}

// func (n *Module) FindNotesBy
