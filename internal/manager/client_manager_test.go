package manager

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cathatino/notification-service/pkg/sql/pg"
	"github.com/stretchr/testify/require"
)

var config *pg.Config

func fetchNewClientManager(t *testing.T) ClientManager {
	conn, err := pg.New(config)
	if err != nil {
		t.Fatal(err)
	}
	return NewClientManager(conn)
}

func TestFindClientById(t *testing.T) {
	clientManager := fetchNewClientManager(t)

	var clientIdOne int64 = 1
	client, err := clientManager.GetClientByClientId(context.Background(), clientIdOne)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, client.ClientId == clientIdOne)
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
	config = &pg.Config{
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
