package eventprocessor

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/domain"
)

type repository interface {
	InsertEvent(ctx context.Context, event domain.Event) error
	SelectUnprocessedEvents(ctx context.Context, limit int64) (events []domain.Event, err error)
	MarkProcessedEvent(ctx context.Context, event domain.Event) error
	UpsertProgress(ctx context.Context, progress domain.Progress) error
	UpsertRating(ctx context.Context, rating domain.Rating) error
}
