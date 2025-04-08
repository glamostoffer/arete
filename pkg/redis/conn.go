package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Connector struct {
	*redis.Client
	cfg Config
}

const (
	componentName = "redis"
)

func New(cfg Config) Connector {
	conn := Connector{
		cfg: cfg,
	}

	conn.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return conn
}

func (c *Connector) Start(ctx context.Context) error {
	_, err := c.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	return c.Close()
}

func (c *Connector) GetName() string {
	return componentName
}
