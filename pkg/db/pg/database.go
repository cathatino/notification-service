package pg

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	dialectsLock sync.RWMutex
	dials        map[string]*Database
)

type Database struct {
	pool *pgxpool.Pool
	dsn  string
	// TODO handle Database & Config
}

func (db *Database) New(dsn string) (*Database, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &Database{
		pool: pool,
		dsn:  dsn,
	}, nil
}

func (db *Database) GetDatabase() (*Database, error) {
	dialectsLock.Lock()
	defer dialectsLock.Unlock()

	if db.dsn == "" {
		return nil, ErrorInvalidDsnString
	}
	if dials == nil {
		dials = make(map[string]*Database)
	}

	database, ok := dials[db.dsn]
	if ok {
		return database, nil
	}

	database, err := db.New(db.dsn)
	if err != nil {
		return nil, err
	}
	dials[db.dsn] = database
	return database, nil
}
