package producer

import (
	"context"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

const (
	componentName = "kafka-producer"
)

type Producer struct {
	cfg Config

	*kafka.Writer
}

func New(cfg Config) Producer {
	return Producer{
		cfg: cfg,
	}
}

func (p *Producer) Start(ctx context.Context) error {
	p.Writer = &kafka.Writer{
		Addr:     kafka.TCP(p.cfg.Address),
		Topic:    p.cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return nil
}

func (p *Producer) Stop(ctx context.Context) error {
	if err := p.Close(); err != nil {
		return errors.Wrap(err, "failed to close writer")
	}

	return nil
}

func (p *Producer) GetName() string {
	return componentName
}
