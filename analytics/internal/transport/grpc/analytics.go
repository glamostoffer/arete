package grpc

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/service/dto"
	v1 "github.com/glamostoffer/arete/analytics/pkg/api/grpc/v1"
	"github.com/gofrs/uuid"
)

type handler struct {
	v1.UnimplementedAnalyticsServer

	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetUserRating(
	ctx context.Context,
	req *v1.GetUserRatingRequest,
) (res *v1.GetUserRatingResponse, err error) {
	courseID, err := uuid.FromString(req.GetCourseID())
	if err != nil {
		return nil, err
	}

	out, err := h.service.GetRating(ctx, dto.GetRatingRequest{
		UserID:   req.GetUserID(),
		CourseID: courseID,
	})
	if err != nil {
		return nil, err
	}

	return out.ToProto(), nil
}

func (h *handler) GetUserProgress(
	ctx context.Context,
	req *v1.GetUserProgressRequest,
) (res *v1.GetUserProgressResponse, err error) {
	courseID, err := uuid.FromString(req.GetCourseID())
	if err != nil {
		return nil, err
	}

	out, err := h.service.GetProgress(ctx, dto.GetProgressRequest{
		UserID:   req.GetUserID(),
		CourseID: courseID,
	})
	if err != nil {
		return nil, err
	}

	return out.ToProto(), nil
}
