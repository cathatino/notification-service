package sql

import (
	"github.com/jmoiron/sqlx"
)

// A basic inteface for Database Access
type Connector interface {
	GetDB() (*sqlx.DB, error)
}
