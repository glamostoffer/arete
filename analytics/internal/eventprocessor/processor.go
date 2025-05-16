package eventprocessor

import (
	"context"
	"errors"

	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/gofiber/fiber/v2/log"
	"github.com/segmentio/kafka-go"
)

type processor struct {
	cfg      Config
	consumer consumer
	repo     repository
}

func New(cfg Config, consumer consumer, repo repository) *processor {
	return &processor{
		cfg:      cfg,
		consumer: consumer,
		repo:     repo,
	}
}

func (p *processor) Start(ctx context.Context) error {
	go func() {
		for {
			msg, err := p.consumer.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}

				log.Error("error reading message", err)
				continue
			}

			if err := p.SaveEvent(ctx, msg); err != nil {
				log.Error("error saving event", err)
			}
		}
	}()

	return nil
}

func (p *processor) SaveEvent(ctx context.Context, msg kafka.Message) error {
	event := domain.EventFromMessage(msg)

	return p.repo.SaveEvent(ctx, event)
}

func (p *processor) Stop(_ context.Context) error {

	return nil
}
