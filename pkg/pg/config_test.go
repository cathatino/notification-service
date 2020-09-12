package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPsqlInfo(t *testing.T) {
	config := &Config{
		Host:     "host",
		Port:     "5432",
		User:     "user",
		Password: "password",
		Dbname:   "dbname",
	}
	require.Equal(
		t,
		"host=host port=5432 user=user password=password dbname=dbname sslmode=disable",
		config.GetPsqlInfo(),
	)

}
