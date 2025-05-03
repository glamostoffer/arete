package grpc

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

func (h *handler) GetCourseLessons(ctx context.Context, req *v1.GetCourseLessonsRequest) (res *v1.GetCourseLessonsResponse, err error) {
	out, err := h.service.GetCourseLessons(ctx, dto.GetCourseLessonsRequest{
		CourseID: req.GetCourseID(),
		Limit:    req.GetLimit(),
		Offset:   req.GetOffset(),
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) GetLessonDetails(ctx context.Context, req *v1.GetLessonDetailsRequest) (res *v1.GetLessonDetailsResponse, err error) {
	out, err := h.service.GetLessonDetails(ctx, dto.GetLessonDetailsRequest{
		LessonID: req.GetLessonID(),
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}
