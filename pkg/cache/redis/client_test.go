package redis

import (
	"os"
	"testing"
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
	config = &Config{
		Host:        env("REDIS_TEST_HOST", "host"),
		Password:    env("REDIS_TEST_PASSWORD", "password"),
		IdleTimeout: 300 * time.Second,
		MaxActive:   100,
		MaxIdle:     10,
	}
}

func TestGetNewRedisClient(t *testing.T) {
	_, err := GetNewRedisClient(config)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisClientPing(t *testing.T) {
	connector, err := GetNewRedisClient(config)
	if err != nil {
		t.Fatal(err)
	}

	if err = connector.Ping(); err != nil {
		t.Fatal(err)
	}
}
