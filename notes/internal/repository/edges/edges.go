package edges

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sosisterrapstar/simple_rest_go/internal/core/edges"
)

type EdgeRepo struct {
	conn *pgxpool.Conn
}

func New(c *pgxpool.Conn) *EdgeRepo {
	return &EdgeRepo{
		conn: c,
	}
}

func (er *EdgeRepo) GetNotesGraph(ctx context.Context, userId uuid.UUID) (*edges.Graph, error) {
}

func (er *EdgeRepo) NewEdge(ctx context.Context, start uuid.UUID, end uuid.UUID, direction edges.Direction) {

}
