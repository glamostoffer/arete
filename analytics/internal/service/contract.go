package service

import (
	"context"
	"time"

	"github.com/glamostoffer/arete/analytics/internal/domain"
)

type repository interface {
	GetUserProgress(ctx context.Context, userID int64) (domain.UserCourseProgress, error)
	GetUserRating(ctx context.Context, userID int64) (domain.CourseRating, error)
	GetUserGlobalRating(ctx context.Context, userID int64) (domain.GlobalRating, error)
	GetUserStats(ctx context.Context, userID int64) (domain.UserStats, error)
	UpdateUserStats(ctx context.Context, stats domain.UserStats) error
	SelectAndLockEvent(ctx context.Context, maxAttempts int64, lockTime time.Duration) (domain.Event, error)
	UpdateEvent(ctx context.Context, event domain.Event) error
}
