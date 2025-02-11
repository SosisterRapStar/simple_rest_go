package postgres

import (
	"context"
	"first-proj/appconfig"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// postgres db adapter
type postgres struct {
	pool *pgxpool.Pool
}

func newPostgres() *postgres{
	config, err := pgxpool.ParseConfig(appconfig.Config.DB.Url)
	if err != nil {
		fmt.Println("Error during pgx config creating")
		os.Exit(1)
	}
	config.MaxConns = appconfig.Config.DB.PGXConfig.MaxConns
	config.MinConns = appconfig.Config.DB.PGXConfig.MinConns
	connPool, err := pgxpool.NewWithConfig(context.Background(), config)r
	return &postgres{pool: connPool}

}

func (p *postgres) Close() {
	p.pool.Close()
}