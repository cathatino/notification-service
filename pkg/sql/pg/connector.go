package pg

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	psqlDriverName string = "postgres"
)

var (
	dialectsLock sync.RWMutex // TODO check locking mechanism
	dials        map[string]*sqlx.DB
)

// Connector is exposed to make the db connection more easier
// By abstracting the db & config property
type Connector struct {
	db     *sqlx.DB
	config *Config
}

// Open New Connection
// Initialize New Connector
func New(config *Config) (*Connector, error) {
	db, err := sqlx.Open(psqlDriverName, config.GetPsqlInfo())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	return &Connector{
		db:     db,
		config: config,
	}, err
}

// GetDB
// returns an existing or returns new Connector
func (c *Connector) GetDB() (*sqlx.DB, error) {
	dialectsLock.Lock()
	defer dialectsLock.Unlock()
	if dials == nil {
		dials = make(map[string]*sqlx.DB)
	}
	psqlInfo := c.config.GetPsqlInfo()
	db, ok := dials[psqlInfo]
	if ok {
		return db, nil
	}

	connector, err := New(c.config)
	if err != nil {
		return nil, err
	}
	dials[psqlInfo] = connector.db
	return connector.db, nil
}

// TODO have a close func
