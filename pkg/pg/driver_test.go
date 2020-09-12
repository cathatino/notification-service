package pg

import (
	"os"
	"time"
)

var config *Config

func init() {
	// get env variables
	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}

	// init config
	config = &Config{
		Host:            env("PSQL_TEST_HOST", "host"),
		Port:            env("PSQL_TEST_PORT", "5432"),
		User:            env("PSQL_TEST_USER", "user"),
		Password:        env("PSQL_TEST_PASSWORD", "password"),
		Dbname:          env("PSQL_TEST_DBNAME", "dbname"),
		MaxOpenConns:    10,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Hour,
	}
}
