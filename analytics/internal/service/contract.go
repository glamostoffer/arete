package service

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/gofrs/uuid"
)

type repository interface {
	GetRating(ctx context.Context, userID int64, courseID uuid.UUID) (rating domain.Rating, err error)
	GetProgress(ctx context.Context, userID int64, courseID uuid.UUID) (progress domain.Progress, err error)
}
