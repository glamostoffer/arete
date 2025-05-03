package service

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
)

func (s *service) GetCourseLessons(ctx context.Context, req dto.GetCourseLessonsRequest) (res dto.GetCourseLessonsResponse, err error) {
	lessons, err := s.repo.GetLessons(ctx, req.CourseID, req.Limit, req.Offset)
	if err != nil {
		return res, err
	}

	if len(lessons) > int(req.Limit) {
		res.HasNext = true
	}
	res.Lessons = lessons

	return res, nil
}

func (s *service) GetLessonDetails(ctx context.Context, req dto.GetLessonDetailsRequest) (res dto.GetLessonDetailsResponse, err error) {
	lesson, err := s.repo.GetLessonDetails(ctx, req.LessonID)
	if err != nil {
		return res, err
	}

	res.Lesson = lesson
	return res, err
}
