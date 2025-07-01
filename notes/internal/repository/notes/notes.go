package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sosisterrapstar/simple_rest_go/internal/core/notes"
)

type NoteRepo struct {
	conn pgxpool.Pool
}

func (nr *NoteRepo) Save(ctx context.Context, params *notes.CreateNote) (*notes.SavedNote, error) {
	var (
		sn notes.SavedNote
	)
	if err := nr.conn.QueryRow(ctx, createNoteQuery, params.Name, params.Content, params.ExpiresAt).Scan(&sn.ID,
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
	if _, err := nr.conn.Exec(ctx, deleteNoteQuery); err != nil {
		return err
	}
	return nil
}

// not implemented
func (nr *NoteRepo) GetUsersNotes(ctx context.Context, userId uuid.UUID) ([]*notes.Note, error) {
	rows, err := nr.conn.Query(ctx, getUserNotesQuery, userId)
	if err != nil {
		return nil, err
	}

	noteModels, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*notes.Note, error){
		var (
			n notes.Note
			tags string = ""
		)
		if err := row.Scan(
			&n.ID,
			&n.Name,
			&n.Content,
			&n.ExpiresAt,
		)
	})
	if err != nil {
		return nil, err
	}

}

func (nr *NoteRepo) fromPostgresArrayToGo(pArr string) []string {

}

// TODO: think about what to do with query (it's actually shoould be done using elastic)
// func (nr *NoteRepo) List(ctx context.Context, userId uuid.UUID, paging core.LimitOffsetPaging) {
// 	limit := paging.Limit
// 	offset := paging.Offset

// }

func (nr *NoteRepo) Update(ctx context.Context, id uuid.UUID, params *notes.UpdateNote) (*notes.SavedNote, error) {
	var (
		un notes.SavedNote
	)
	if err := nr.conn.QueryRow(ctx,
		updateQueryNote,
		id,
		params.Content,
		params.Name,
		params.ExpiresAt,
	).Scan(un.ID, un.Name, un.Content, un.ExpiresAt, un.CreatedAt); err != nil {
		return nil, err
	}
	return &un, nil
}
