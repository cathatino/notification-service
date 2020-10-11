package connector

import "github.com/jmoiron/sqlx"

type Connector interface {
	GetDB() (*sqlx.DB, error)
}
