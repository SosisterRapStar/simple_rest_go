package connection

import (
	"context"
	"first-proj/appconfig"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// there is no any abstraction on pgx pool so
// there is no any Open/Close interface methods, so the dependency provider should close the pool "manualy"
// using pgx functions intself after Service work was done
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
