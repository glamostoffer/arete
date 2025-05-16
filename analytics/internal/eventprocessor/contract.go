package eventprocessor

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/segmentio/kafka-go"
)

type repository interface {
	SaveEvent(ctx context.Context, event domain.Event) error
	// SelectUnprocessedEvents(ctx context.Context, limit int64) (events []domain.Event, err error)
	// MarkProcessedEvent(ctx context.Context, event domain.Event) error
	// UpsertProgress(ctx context.Context, progress domain.Progress) error
	// UpsertRating(ctx context.Context, rating domain.Rating) error
}

type consumer interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}
