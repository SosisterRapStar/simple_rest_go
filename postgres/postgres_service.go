package postgres

import (
	"context"
	"errors"
	"first-proj/domain"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(db *pgxpool.Pool) *PostgresService {
	return &PostgresService{db: db}
}

type PostgresService struct {
	db *pgxpool.Pool
}

func (pgs *PostgresService) CreateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return uuid.Nil, err

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.NewV7()

	if err != nil {
		return uuid.Nil, err
	}

	note.Id = id
	query := `INSERT INTO notes.note (id, title, content) VALUES ($1, $2, $3)`

	_, err = tx.Exec(ctx, query, note.Id, note.Title, note.Content)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error occurred during note insertion: %w", err)
	}
	defer tx.Rollback(ctx)

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return note.Id, nil

}

// vot bi suda contextniy manager
func (pgs *PostgresService) GetNote(ctx context.Context, id string) (*domain.Note, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	note := &domain.Note{}
	note_id, _ := uuid.Parse(id)

	err = conn.QueryRow(ctx, "SELECT id, title, content FROM notes.note WHERE id = $1", note_id).Scan(&note.Id, &note.Title, &note.Content)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}

	return note, nil
}

func (pgs *PostgresService) DeleteNote(ctx context.Context, id string) (uuid.UUID, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return uuid.Nil, err

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return uuid.Nil, err
	}
	note_id, _ := uuid.Parse(id)
	_, err = tx.Exec(ctx, "DELETE FROM notes.note WHERE id = $1", note_id)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return note_id, nil

}

func (pgs *PostgresService) updateOldObj(old *domain.Note, upd *domain.UpdateNote) *domain.Note {
	updValue := reflect.ValueOf(upd).Elem()
	oldValue := reflect.ValueOf(old).Elem()

	for i := 0; i < updValue.NumField(); i++ {
		oldField := oldValue.Field(i + 1)
		updField := updValue.Field(i)
		if !updField.IsNil() {
			updFieldValue := updField.Elem()
			oldField.Set(updFieldValue)
		}
	}
	return old
}

func (pgs *PostgresService) UpdateNote(ctx context.Context, upd *domain.UpdateNote, id string) (*domain.Note, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, err

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, err
	}
	old_note, err := pgs.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	updated_value := pgs.updateOldObj(old_note, upd)
	_, err = tx.Exec(ctx, "UPDATE notes.note SET title = $1, content = $2 WHERE id = $3", updated_value.Title, updated_value.Content, updated_value.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}
	defer tx.Rollback(ctx)

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return updated_value, nil
}
