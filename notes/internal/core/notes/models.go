package notes

import (
	"time"

	"github.com/google/uuid"
)

type Direction uint8
type Tag string

var (
	Indirect Direction = 0
	Direct   Direction = 1
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
	Tags []string `json:"tags,omitempty"`
}

type SavedNote struct {
	NoteBase
}

type CreateNote struct {
	Name      uuid.UUID `json:"name"`
	Content   uuid.UUID `json:"content"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type UpdateNote struct {
	Name    uuid.UUID `json:"name"`
	Content uuid.UUID `json:"content"`
	Tags    []string  `json:"tags"`
}

type Edge struct {
	EdgeType  string    `json:"edge_type"`
	Direction Direction `json:"direction"`
	Start     *Note     `json:"start"`
	End       *Note     `json:"end"`
}

type Graph []*Edge

type Query struct {
	QueryStr string `json:"query_str,omitempty"`
	Tags     []Tag  `json:"tags,omitempty"`
}
