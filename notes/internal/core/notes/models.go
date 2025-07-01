package notes

import (
	"time"

	"github.com/google/uuid"
)

type State string

var (
	Done    State = "done"
	Current State = "current"
	Empty   State = "empty"
)

type NoteBase struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Name      uuid.UUID `json:"name"`
	Content   uuid.UUID `json:"content"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	NoteBase
}

type CreateNote struct {
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type UpdateNote struct {
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Query struct {
	QueryStr string   `json:"query_str,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}
