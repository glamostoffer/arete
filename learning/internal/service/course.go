package service

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
)

func (s *service) GetCoursesList(ctx context.Context, req dto.GetCoursesListRequest) (res dto.GetCoursesListResponse, err error) {
	courses, err := s.repo.GetCourses(ctx, req.Categories, req.UserID, req.Limit, req.Offset)
	if err != nil {
		return res, err
	}

	if len(courses) > int(req.Limit) {
		res.HasNext = true
	}
	res.Courses = courses

	return res, nil
}

func (s *service) GetCourseCategories(ctx context.Context, req dto.GetCourseCategoriesRequest) (res dto.GetCourseCategoriesResponse, err error) {
	categories, err := s.repo.GetCourseCategories(ctx)
	if err != nil {
		return res, err
	}

	res.Categories = categories
	return res, nil
}

func (s *service) EnrollToCourse(ctx context.Context, req dto.EnrollToCourseRequest) (res dto.EnrollToCourseResponse, err error) {
	err = s.repo.EnrollUserToCourse(ctx, req.UserID, req.CourseID)
	if err != nil {
		return res, err
	}

	return res, nil
}
