package repository

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/domain"
)

func (r *repository) GetLessons(ctx context.Context, courseID int64, limit, offset int64) (lessons []domain.Lesson, err error) {
	err = r.db.SelectContext(
		ctx,
		&lessons,
		queryGetLessons,
		courseID,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (r *repository) GetLessonDetails(ctx context.Context, lessonID int64) (lesson domain.Lesson, err error) {
	err = r.db.GetContext(
		ctx,
		&lesson,
		queryGetLessonDetails,
		lessonID,
	)
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}
