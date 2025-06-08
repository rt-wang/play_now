package queue

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
)

// Client handles the connection to NATS JetStream.
type Client struct {
	js jetstream.JetStream
}

// New creates a new NATS JetStream client.
func New(ctx context.Context, url string) (*Client, error) {
	nc, err := jetstream.Connect(url)
	if err != nil {
		return nil, err
	}
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}
	return &Client{js: js}, nil
}
