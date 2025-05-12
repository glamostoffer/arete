package consumer

import (
	"context"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

const (
	componentName = "kafka-consumer"
)

type Consumer struct {
	cfg Config

	*kafka.Reader
}

func New(cfg Config) Consumer {
	return Consumer{
		cfg: cfg,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   c.cfg.Brokers,
		Topic:     c.cfg.Topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	c.SetOffset(kafka.LastOffset)

	return nil
}

func (c *Consumer) Stop(ctx context.Context) error {
	if err := c.Close(); err != nil {
		return errors.Wrap(err, "failed to close reader")
	}

	return nil
}

func (c *Consumer) GetName() string {
	return componentName
}
