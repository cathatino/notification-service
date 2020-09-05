package pg

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	connString := os.Getenv("PG_CI_CD_DATABASE")
	pool, err := pgxpool.Connect(context.Background(), connString)
	defer pool.Close()
	require.NoError(t, err)
	assert.Equal(t, connString, pool.Config().ConnString())
}

func TestQuery(t *testing.T) {
	t.Parallel()
	connString := os.Getenv("PG_CI_CD_DATABASE")
	pool, err := pgxpool.Connect(context.Background(), connString)
	defer pool.Close()
	require.NoError(t, err)

	const rowsLength = 10
	rows, err := pool.Query(context.Background(), "select generate_series(1,$1)", rowsLength)
	require.NoError(t, err)

	stats := pool.Stat()
	assert.EqualValues(t, 1, stats.AcquiredConns())
	assert.EqualValues(t, 1, stats.TotalConns())

	rowsCount := 0
	for rows.Next() {
		rowsCount += 1
	}
	assert.Equal(t, rowsCount, rowsLength)
	assert.NoError(t, rows.Err())
}
