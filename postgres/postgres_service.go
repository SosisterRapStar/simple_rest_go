package postgres

import (
	"context"
	"errors"
	"first-proj/domain"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresService struct {
	db *pgxpool.Pool
}

func (pgs *PostgresService) CreateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return err

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	id, err := uuid.NewV7()
	
	if err != nil {
		return err
	}

	note.Id = id
	query := `INSERT INTO notes (id, title, content) VALUES ($1, $2, $3)`

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

	if err != nil {
		return nil, err
	}

 	
	if err != nil {
		return nil, err
	}

	note = &domain.Note{}
	err := conn.QueryRow(ctx, "SELECT id, title, content FROM notes WHERE id = $1", note_id).Scan(&node.Id, &note.Title, &note.Content)
	
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		} 
		return nil, err
	}

	return note, nil
}

func (pgs *PostgresService) DeleteNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return err

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}
	
}

func (pgs *PostgresService) UpdateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}



type DeadORM