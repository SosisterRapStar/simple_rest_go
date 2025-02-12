package postgres

import (
	"context"
	"first-proj/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresService struct {
	db *pgxpool.Pool
}

func (pgs *PostgresService) CreateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {
}

func (pgs *PostgresService) FindNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}

func (pgs *PostgresService) DeleteNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}

func (pgs *PostgresService) UpdateNote(ctx context.Context, note *domain.Note) (uuid.UUID, error) {

}
