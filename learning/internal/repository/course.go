package repository

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/domain"
	"github.com/lib/pq"
)

func (r *repository) GetCourses(ctx context.Context, categories []string, userID, limit, offset int64) (courses []domain.Course, err error) {
	err = r.db.SelectContext(
		ctx,
		&courses,
		queryGetCourses,
		userID,
		pq.Array(categories),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *repository) GetCourseCategories(ctx context.Context) (categories []string, err error) {
	err = r.db.SelectContext(
		ctx,
		&categories,
		queryGetCourseCategories,
	)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
