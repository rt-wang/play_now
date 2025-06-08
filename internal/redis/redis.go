package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client wraps the redis client used throughout the service.
type Client struct {
	*redis.Client
}

// New creates a new redis client instance.
func New(addr string) *Client {
	r := redis.NewClient(&redis.Options{Addr: addr})
	return &Client{Client: r}
}

// Ping checks connectivity with redis.
func (c *Client) Ping(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}
