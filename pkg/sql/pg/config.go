package pg

import (
	"fmt"
	"time"
)

const (
	psqlFmt string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

type Config struct {
	Host            string
	Port            string
	User            string
	Password        string
	Dbname          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func (c *Config) GetPsqlInfo() string {
	return fmt.Sprintf(
		psqlFmt,
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Dbname,
	)
}
