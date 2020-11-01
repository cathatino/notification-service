package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	rd "github.com/garyburd/redigo/redis"
)

const (
	CommandPing string = "PING"
	ServerTcp   string = "tcp"
)

type Connector struct {
	pool *rd.Pool
}

func (c *Connector) Ping() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return fmt.Errorf("can't 'PING' redis: %v", err)
	}
	return nil
}

func (c *Connector) Get(key string) ([]byte, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func (c *Connector) Set(key string, value []byte) error {
	conn := c.pool.Get()
	defer conn.Close()
    _, err := conn.Do("SET", key, value)
    return err
}

type Config struct {
	// Maximum number of idle connections in the pool.
	MaxIdle int

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// Redis Host to be connect to.
	Host string

	// Password of redis service.
	Password string
}

func GetNewRedisClient(config *Config) (*Connector, error) {
	return &Connector{
		pool: &rd.Pool{
			MaxIdle:     config.MaxIdle,
			IdleTimeout: config.IdleTimeout,

			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(ServerTcp,
					config.Host,
					rd.DialPassword(config.Password),
				)
				if err != nil {
					return nil, err
				}
				return c, err
			},

			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do(CommandPing)
				return err
			},
		},
	}, nil
}
