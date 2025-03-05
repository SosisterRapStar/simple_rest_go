package connections

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// there is no any abstraction on pgx pool so
// there is no any Open/Close interface methods, so the dependency provider should close the pool "manualy"
// using pgx functions intself after Service work was done
func NewPool(maxCons int32, minCons int32, url string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("Error during pgx config creating")
		os.Exit(1)
	}
	config.MaxConns = int32(maxCons)
	config.MinConns = int32(minCons)
	connPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error occured during pool creating")
		os.Exit(1)
	}
	return connPool

}
