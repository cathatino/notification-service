package pg

import (
	"os"
	"testing"
	"time"
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
	connector, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	// test for each execution testing
	connectorTest := &ConnectorTest{t, connector}
	for _, execTesting := range execTestings {
		execTesting(connectorTest)
	}
}

func TestConnectorConnection(t *testing.T) {
	connectionTesting := func(ct *ConnectorTest) {
		_, err := ct.connector.Open()
		if err != nil {
			ct.Fatal(err)
		}
	}
	runTests(
		t,
		config,
		connectionTesting,
	)
}
