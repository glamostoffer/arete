package dto

import (
	"github.com/glamostoffer/arete/analytics/internal/domain"
	v1 "github.com/glamostoffer/arete/analytics/pkg/api/grpc/v1"
	"github.com/gofrs/uuid"
)

type GetRatingRequest struct {
	UserID   int64
	CourseID uuid.UUID
}
type GetRatingResponse struct {
	Rating domain.Rating
}

func (r *GetRatingResponse) ToProto() *v1.GetUserRatingResponse {
	return &v1.GetUserRatingResponse{}
}

type GetProgressRequest struct {
	UserID   int64
	CourseID uuid.UUID
}
type GetProgressResponse struct {
	Progress domain.Progress
}

func (r *GetProgressResponse) ToProto() *v1.GetUserProgressResponse {
	return &v1.GetUserProgressResponse{}
}
