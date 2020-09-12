package driver

import "database/sql"

type DB struct {
	*sql.DB
}
