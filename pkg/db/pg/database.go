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
	if db.dsn == "" {
		return nil, ErrorInvalidDsnString
	}
	if dials == nil {
		initDials()
	}

	database, ok := dials[db.dsn]
	if ok {
		return database, nil
	}

	database, err := db.New(db.dsn)
	if err != nil {
		return nil, err
	}

	updateDails(db.dsn, database)
	return database, nil
}

func initDials() {
	dialectsLock.Lock()
	defer dialectsLock.Unlock()
	dials = make(map[string]*Database)
}

func updateDails(dsn string, db *Database) {
	dialectsLock.Lock()
	defer dialectsLock.Unlock()
	dials[dsn] = db
}
