package edges

import "github.com/sosisterrapstar/simple_rest_go/internal/core/notes"

type Direction uint8
type EdgeType string

var (
	Inderect Direction = 0
	Direct   Direction = 1
)

type Edge struct {
	EdgeType  string      `json:"edge_type"`
	Direction Direction   `json:"direction"`
	Start     *notes.Note `json:"start"`
	End       *notes.Note `json:"end"`
}

type Graph []*Edge
