package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sosisterrapstar/simple_rest_go/internal/core"
	"github.com/sosisterrapstar/simple_rest_go/internal/core/notes"
)

type NoteRepo struct {
	conn pgxpool.Pool
}

func (nr *NoteRepo) Save(ctx context.Context, params *notes.CreateNote) (*notes.SavedNote, error) {
	var (
		sn notes.SavedNote
	)
	if err := n.conn.QueryRow(ctx, createNoteQuery, params.Name, params.Content, params.ExpiresAt).Scan(&sn.ID,
		&sn.Name,
		&sn.Content,
		&sn.ExpiresAt,
		&sn.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &sn, nil
}

func (nr *NoteRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := n.conn.Exec(ctx, deleteNoteQuery); err != nil {
		return err
	}
	return nil
}

func (nr *NoteRepo) Get(ctx context.Context, userId uuid.UUID, id uuid.UUID) (*notes.Note, error) {
}

func (nr *NoteRepo) List(ctx context.Context, userId uuid.UUID, paging core.LimitOffsetPaging) {

}

func (nr *NoteRepo) Update() {

}
