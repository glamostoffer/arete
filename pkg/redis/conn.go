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
	return Connector{
		cfg: cfg,
	}
}

func (c *Connector) Start(ctx context.Context) error {
	c.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.cfg.Host, c.cfg.Port),
		Password: c.cfg.Password,
		DB:       c.cfg.DB,
	})

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
