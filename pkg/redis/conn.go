package redis

import (
	"context"
	"fmt"

	"github.com/glamostoffer/arete/pkg/component"
	"github.com/redis/go-redis/v9"
)

type connector struct {
	*redis.Client
	cfg Config
}

const (
	componentName = "redis"
)

func New(cfg Config) component.Component {
	return &connector{
		cfg: cfg,
	}
}

func (c *connector) Start(ctx context.Context) error {
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

func (c *connector) Stop(ctx context.Context) error {
	return c.Close()
}

func (c *connector) GetName() string {
	return componentName
}
