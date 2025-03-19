package pg

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/glamostoffer/arete/analytics/pkg/errlist"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UpsertRating(ctx context.Context, rating domain.Rating) error {
	res, err := r.db.ExecContext(
		ctx,
		queryUpsertRating,
		rating.ID,
		rating.UserID,
		rating.CourseID,
		rating.Score,
		rating.LastUpdated,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errlist.ErrInvalidAffectedRows
	}

	return nil
}

func (r *repository) GetRating(ctx context.Context, userID int64, courseID uuid.UUID) (rating domain.Rating, err error) {
	err = r.db.GetContext(
		ctx,
		&rating,
		queryGetRating,
		userID,
		courseID,
	)
	if err != nil {
		return domain.Rating{}, err
	}

	return rating, nil
}

func (r *repository) UpsertProgress(ctx context.Context, progress domain.Progress) error {
	res, err := r.db.ExecContext(
		ctx,
		queryUpsertProgress,
		progress.ID,
		progress.UserID,
		progress.CourseID,
		progress.TotalMaterialsCompleted,
		progress.TotalTasksCompleted,
		progress.TotalQuizzesCompleted,
		progress.TotalScore,
		progress.CompletionRate,
		progress.LastUpdated,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errlist.ErrInvalidAffectedRows
	}

	return nil
}

func (r *repository) GetProgress(ctx context.Context, userID int64, courseID uuid.UUID) (progress domain.Progress, err error) {
	err = r.db.GetContext(
		ctx,
		&progress,
		queryGetProgress,
		userID,
		courseID,
	)
	if err != nil {
		return domain.Progress{}, err
	}

	return progress, nil
}

func (r *repository) InsertEvent(ctx context.Context, event domain.Event) error {
	res, err := r.db.ExecContext(
		ctx,
		queryInsertEvent,
		event.ID,
		event.UserID,
		event.CourseID,
		event.Type,
		event.Data,
		event.CreatedAt,
		event.ProcessedAt,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errlist.ErrInvalidAffectedRows
	}

	return nil
}

func (r *repository) SelectUnprocessedEvents(ctx context.Context, limit int64) (events []domain.Event, err error) {
	err = r.db.SelectContext(
		ctx,
		&events,
		querySelectUnprocessedEvents,
		limit,
	)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *repository) MarkProcessedEvent(ctx context.Context, event domain.Event) error {
	res, err := r.db.ExecContext(
		ctx,
		queryUpdateEvent,
		event.ID,
		event.ProcessedAt,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errlist.ErrInvalidAffectedRows
	}

	return nil
}
