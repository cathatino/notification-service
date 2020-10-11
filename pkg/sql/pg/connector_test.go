package pg

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	config    *Config
	available bool
)

type ConnectorTest struct {
	*testing.T
	connector *Connector
}

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

// runTests establishs DB connection first
// then start with each execution testing
func runTests(
	t *testing.T,
	config *Config,
	execTestings ...func(*ConnectorTest),
) {
	// establish connection
	conn, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	// test for each execution testing
	connectorTest := &ConnectorTest{t, conn}
	for _, execTesting := range execTestings {
		execTesting(connectorTest)
	}
}

// TestConnectorConnection tests for the db connection from the connector's connection pool
func TestConnectorConnection(t *testing.T) {
	connectionTesting := func(ct *ConnectorTest) {
		_, err := ct.connector.GetDB()
		if err != nil {
			ct.Fatal(err)
		}
	}
	runTests(t, config, connectionTesting)
}

// TestConnectorRunDBQuery tests for basic db query,
// including "select generate_series"
func TestConnectorRunDBQuery(t *testing.T) {
	connectionDBQueryTesting := func(ct *ConnectorTest) {
		db, err := ct.connector.GetDB()
		if err != nil {
			ct.Fatal(err)
		}
		rows, err := db.Query("select generate_series(1, 10)")
		if err != nil {
			ct.Fatal(err)
		}

		counter := 0
		for rows.Next() {
			counter++
			var cur int
			err := rows.Scan(&cur)
			if err != nil {
				ct.Fatal(err)
			}
			assert.Equal(t, cur, counter)
		}
	}
	runTests(t, config, connectionDBQueryTesting)
}
