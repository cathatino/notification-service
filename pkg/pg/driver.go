package pg

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

const (
	psqlDriverName string = "postgres"
)

var (
	dialectsLock sync.RWMutex // TODO check locking mechanism
	dials        map[string]*Connector
)

// Connector is exposed to make the db connection more easier
// By abstracting the db & config property
type Connector struct {
	db     *sqlx.DB
	config *Config
}

// Open New Connection
// Initialize New Connector
func (c *Connector) New(config *Config) (*Connector, error) {
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

// Get
// returns an existing or returns new Connector
func (c *Connector) Get() (*Connector, error) {
	dialectsLock.Lock()
	defer dialectsLock.Unlock()
	if dials == nil {
		dials = make(map[string]*Connector)
	}
	psqlInfo := c.config.GetPsqlInfo()
	connector, ok := dials[psqlInfo]
	if ok {
		return connector, nil
	}

	connector, err := c.New(c.config)
	if err != nil {
		return nil, err
	}
	dials[psqlInfo] = connector
	return connector, nil
}