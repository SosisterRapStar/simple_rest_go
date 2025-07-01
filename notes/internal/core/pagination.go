package core

import "github.com/sosisterrapstar/simple_rest_go/internal/core/notes"

type LimitOffsetPaging struct {
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Query  notes.Query `json:"query,omitempty"`
}

type CursorPaging struct {
	NextPageToken string
	Limit         int
	Query         notes.Query `json:"query,omitempty"`
}
