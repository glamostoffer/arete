package grpc

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

func (h *handler) GetCoursesList(ctx context.Context, req *v1.GetCoursesListRequest) (res *v1.GetCoursesListResponse, err error) {
	in := dto.GetCoursesListRequest{
		UserID:     req.GetUserID(),
		Categories: nil,
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	}

	if len(req.GetCategories()) != 0 {
		in.Categories = req.GetCategories()
	}

	out, err := h.service.GetCoursesList(ctx, in)
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) GetCourseCategories(ctx context.Context, req *v1.GetCourseCategoriesRequest) (res *v1.GetCourseCategoriesResponse, err error) {
	out, err := h.service.GetCourseCategories(ctx, dto.GetCourseCategoriesRequest{})
	if err != nil {
		return res, err
	}

	return &v1.GetCourseCategoriesResponse{
		Categories: out.Categories,
	}, err
}
