package pg

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO handle Database & Config
type Database struct {
	pool *pgxpool.Pool
}

func (database *Database) Query(
	ctx context.Context,
	statement string,
	args ...interface{},
) (pgx.Rows, error) {
	return database.pool.Query(
		ctx,
		statement,
		args,
	)
}

func (database *Database) Exec(
	ctx context.Context,
	statement string,
	args ...interface{},
) (pgconn.CommandTag, error) {
	return database.pool.Exec(
		ctx,
		statement,
		args,
	)
}
