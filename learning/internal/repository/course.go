package repository

import (
	"context"
	"errors"

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

func (r *repository) EnrollUserToCourse(ctx context.Context, userID, courseID int64) error {
	res, err := r.db.ExecContext(
		ctx,
		queryEnrollUserToCourse,
		userID,
		courseID,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("INVALID_AFFECTED_ROWS")
	}

	return nil
}

func (r *repository) GetUserCourses(ctx context.Context, userID int64) (courses []domain.Course, err error) {
	err = r.db.SelectContext(
		ctx,
		&courses,
		queryGetUserCourses,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
