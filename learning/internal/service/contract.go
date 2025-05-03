package service

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/domain"
)

type repository interface {
	GetCourses(ctx context.Context, categories []string, userID, limit, offset int64) (courses []domain.Course, err error)
	GetCourseCategories(ctx context.Context) (categories []string, err error)

	GetLessons(ctx context.Context, courseID int64, limit, offset int64) (lessons []domain.Lesson, err error)
	GetLessonDetails(ctx context.Context, lessonID int64) (lesson domain.Lesson, err error)
}
