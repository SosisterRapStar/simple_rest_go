package connection

import (
	"context"
	"first-proj/appconfig"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// postgres db adapter

func NewPool() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(appconfig.Config.DB.Url)
	if err != nil {
		fmt.Println("Error during pgx config creating")
		os.Exit(1)
	}
	config.MaxConns = appconfig.Config.DB.PGXConfig.MaxConns
	config.MinConns = appconfig.Config.DB.PGXConfig.MinConns
	connPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error occured during pool creating")
		os.Exit(1)
	}
	return connPool

}

// var db *Postgres = newPostgres()
