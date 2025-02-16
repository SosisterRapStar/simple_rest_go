package postgres

import (
	"context"
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
		return uuid.Nil, fmt.Errorf("error occurred getting connection from pool: %w", err)

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error occured during transaction opening: %w", err)
	}

	tx.Begin(ctx)
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error occured during uuid creating: %w", err)
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
		return uuid.Nil, fmt.Errorf("error occurred during transaction commit: %w", err)
	}

	return note.Id, nil

}

func (pgs *PostgresService) GetNote(ctx context.Context, id string) (*domain.Note, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting connection from pool: %w", err)

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error occured during transaction opening: %w", err)
	}

	tx.Begin(ctx)
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error occured during uuid creating: %w", err)
	}

}

func (pgs *PostgresService) DeleteNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}

func (pgs *PostgresService) UpdateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}



type DeadORM