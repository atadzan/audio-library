package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPoolConn(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Verify db connection
	err = dbPool.Ping(ctx)
	if err != nil {
		return dbPool, err
	}
	return dbPool, nil
}
