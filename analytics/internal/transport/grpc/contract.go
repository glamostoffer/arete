package grpc

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/service/dto"
)

type service interface {
	GetRating(
		ctx context.Context,
		req dto.GetRatingRequest,
	) (res dto.GetRatingResponse, err error)
	GetProgress(
		ctx context.Context,
		req dto.GetProgressRequest,
	) (res dto.GetProgressResponse, err error)
}
