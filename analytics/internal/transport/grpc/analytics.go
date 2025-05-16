package grpc

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/service/dto"
	v1 "github.com/glamostoffer/arete/analytics/pkg/api/grpc/v1"
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

func (h *handler) GetUserStats(
	ctx context.Context,
	req *v1.GetUserStatsRequest,
) (res *v1.GetUserStatsResponse, err error) {
	out, err := h.service.GetUserStats(ctx, dto.GetUserStatsRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return res, nil
	}

	return out.ToProto(), nil
}
