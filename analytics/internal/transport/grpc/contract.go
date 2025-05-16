package grpc

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/service/dto"
)

type service interface {
	GetUserStats(ctx context.Context, req dto.GetUserStatsRequest) (res dto.GetUserStatsResponse, err error)
}
