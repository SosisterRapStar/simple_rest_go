package postgres

import "errors"

var (
	ErrTx           = errors.New("Error during transaction exec")
	ErrNoteNotFound = errors.New("Note was not found")
)
