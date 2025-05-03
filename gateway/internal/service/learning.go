package service

import (
	"context"

	"github.com/glamostoffer/arete/gateway/internal/service/dto"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

func (s *service) GetCourseCategories(ctx context.Context, req dto.GetCourseCategoriesRequest) (res dto.GetCourseCategoriesResponse, err error) {
	out, err := s.learning.GetCourseCategories(ctx, &v1.GetCourseCategoriesRequest{})
	if err != nil {
		return res, err
	}

	res.Categories = out.GetCategories()
	return res, nil
}

func (s *service) GetCourses(ctx context.Context, req dto.GetCoursesRequest) (res dto.GetCoursesResponse, err error) {
	out, err := s.learning.GetCoursesList(ctx, &v1.GetCoursesListRequest{
		UserID:     req.UserID,
		Categories: req.Categories,
		Limit:      req.Limit,
		Offset:     req.Offset,
	})
	if err != nil {
		return res, err
	}

	return res.FromProto(out), nil
}

func (s *service) GetCourseLessons(ctx context.Context, req dto.GetCourseLessonsRequest) (res dto.GetCourseLessonsResponse, err error) {
	out, err := s.learning.GetCourseLessons(ctx, &v1.GetCourseLessonsRequest{
		CourseID: req.CourseID,
		Limit:    req.Limit,
		Offset:   req.Offset,
	})
	if err != nil {
		return res, err
	}

	return res.FromProto(out), err
}

func (s *service) GetLessonDetails(ctx context.Context, req dto.GetLessonDetailsRequest) (res dto.GetLessonDetailsResponse, err error) {
	out, err := s.learning.GetLessonDetails(ctx, &v1.GetLessonDetailsRequest{
		LessonID: req.LessonID,
	})
	if err != nil {
		return res, err
	}

	return res.FromProto(out), err
}
