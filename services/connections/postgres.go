package connections

import (
	"context"
	"first-proj/appconfig"

	"github.com/jackc/pgx/v5/pgxpool"
)

var logger = appconfig.GetLogger()

type PostgresConn struct {
	Pool *pgxpool.Pool
}

func (pg *PostgresConn) Open(config appconfig.Config) error {
	pgconfig, err := pgxpool.ParseConfig(config.Url)
	if err != nil {
		logger.Error("Error occured during connection opening")
		return err
	}
	pgconfig.MaxConns = int32(config.MaxConns)
	pgconfig.MinConns = int32(config.MinConns)
	connPool, err := pgxpool.NewWithConfig(context.Background(), pgconfig)
	if err != nil {
		logger.Error("Error occured during postgresql connection pool opening")
		return err
	}
	pg.Pool = connPool
	return nil
}

func (pg *PostgresConn) Close(ctx context.Context) error {
	logger.Info("Started to close postgresql connection pool")
	// blocking call
	pg.Pool.Close()
	logger.Info("PG pool was closed")
	return nil
}
