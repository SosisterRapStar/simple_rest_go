package postgres

import (
	"context"
	"errors"
	"first-proj/domain"
	"first-proj/services"
	"reflect"

	"first-proj/appconfig"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var logger = appconfig.GetLogger()

// inplementation specific errors
// only postgres knows what's happening
// upper interfaces can not handle it so http will throw badrequest for example
// CLI will exit and so on
func NewPostgres(db *pgxpool.Pool) *PostgresService {
	return &PostgresService{db: db}
}

type PostgresService struct {
	db *pgxpool.Pool
}

func (pgs *PostgresService) CreateNote(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	// var note domain.Note
	// note.Title = crt.Title

	if err := note.Validate(); err != nil {
		return nil, services.NewServiceError(err, err)
	}
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}
	defer tx.Rollback(ctx)

	id, err := uuid.NewV7()

	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}

	note.Id = id.String()
	query := `INSERT INTO notes.note (id, title, content) VALUES ($1, $2, $3)`

	_, err = tx.Exec(ctx, query, note.Id, note.Title, note.Content)
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}

	return note, nil

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
			return nil, services.NewServiceError(services.ErrInternalFailure, err) // logic level error
		}
		return nil, err
	}

	return note, nil
}

func (pgs *PostgresService) DeleteNote(ctx context.Context, id string) (string, error) {
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return "", services.NewServiceError(services.ErrInternalFailure, err)

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return "", services.NewServiceError(services.ErrInternalFailure, err)
	}
	defer tx.Rollback(ctx)

	note_id, _ := uuid.Parse(id)

	_, err = tx.Exec(ctx, "DELETE FROM notes.note WHERE id = $1", note_id)
	if err != nil {
		return "", services.NewServiceError(services.ErrInternalFailure, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return "", services.NewServiceError(services.ErrInternalFailure, err)
	}

	return note_id.String(), nil

}

func (pgs *PostgresService) updateOldObj(old *domain.Note, upd *domain.UpdateNote) *domain.Note {
	// n + 1 без ORM = кайф
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
	if upd.Title != nil {
		if *upd.Title == "" {
			return nil, services.NewServiceError(domain.ErrNoteValidation, domain.ErrNoteValidation)
		}
	}
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)

	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}
	defer tx.Rollback(ctx)

	old_note, err := pgs.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	updated_value := pgs.updateOldObj(old_note, upd)
	_, err = tx.Exec(ctx, "UPDATE notes.note SET title = $1, content = $2 WHERE id = $3", updated_value.Title, updated_value.Content, updated_value.Id)
	if err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, services.NewServiceError(services.ErrInternalFailure, err)
	}

	return updated_value, nil
}

func (pgs *PostgresService) FindNotes(ctx context.Context, filter *domain.PaginateFilter) ([]*domain.Note, int, string, error) {
	logger.Info("StartedPagination")
	conn, err := pgs.db.Acquire(ctx)
	if err != nil {
		return nil, 0, "", services.NewServiceError(services.ErrInternalFailure, err)

	}
	defer conn.Release()

	if *filter.Limit > 100 {
		return nil, 0, "", services.NewServiceError(services.ErrTooManyRowsToFetch, services.ErrTooManyRowsToFetch)
	}

	var paginateQuery string
	var rows pgx.Rows
	var args []interface{}

	if filter.NextPageToken == nil {
		logger.Debug("No token provided")
		paginateQuery = "SELECT id, title, content FROM notes.note FETCH FIRST $1 ROWS ONLY"
		args = []interface{}{*filter.Limit}
	} else {
		logger.Debug("Token provided")
		paginateQuery = "SELECT id, title, content FROM notes.note WHERE id > $1 FETCH NEXT $2 ROWS ONLY"
		token, err := uuid.Parse(*filter.NextPageToken)
		if err != nil {
			return nil, 0, "", services.NewServiceError(services.ErrInternalFailure, err)
		}
		args = []interface{}{token, *filter.Limit}
	}

	rows, err = conn.Query(ctx, paginateQuery, args...)
	if err != nil {
		return nil, 0, "", services.NewServiceError(services.ErrInternalFailure, err)
	}
	var notes []*domain.Note
	defer rows.Close()
	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Content)
		if err != nil {
			return nil, 0, "", services.NewServiceError(services.ErrInternalFailure, err)
		}
		notes = append(notes, &note)
	}
	err = rows.Err()
	if err != nil {
		return nil, 0, "", services.NewServiceError(services.ErrInternalFailure, err)
	}
	var nextToken string
	if len(notes) > 0 {
		nextToken = notes[len(notes)-1].Id
	} else {
		nextToken = ""
	}
	return notes, len(notes), nextToken, nil

}
