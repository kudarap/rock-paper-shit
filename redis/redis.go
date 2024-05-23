package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client represents the client for the internal redis package
type Client struct {
	*redis.Client

	config Config
	logger *slog.Logger
}

// NewClient create new instance of redis client.
func NewClient(conf Config, logger *slog.Logger) (*Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr: conf.Addr,
	})
	c := &Client{Client: rc, config: conf, logger: logger}
	if _, err := c.Ping(); err != nil {
		return nil, err
	}
	return c, nil
}

// Ping checks redis connection
func (c *Client) Ping() (ok bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err = c.Client.Ping(ctx).Err(); err != nil {
		return false, fmt.Errorf("redis: could not ping client: %s", err)
	}
	return true, nil
}

// Close closes all connection.
func (c *Client) Close() error {
	return c.Client.Close()
}

// Config represents the redis configuration
type Config struct {
	Addr string
}
