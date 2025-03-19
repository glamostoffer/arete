package service

import (
	"context"

	"github.com/glamostoffer/arete/analytics/internal/service/dto"
)

type service struct {
	repo repository
}

func New(repo repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetRating(
	ctx context.Context,
	req dto.GetRatingRequest,
) (res dto.GetRatingResponse, err error) {
	res.Rating, err = s.repo.GetRating(ctx, req.UserID, req.CourseID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *service) GetProgress(
	ctx context.Context,
	req dto.GetProgressRequest,
) (res dto.GetProgressResponse, err error) {
	res.Progress, err = s.repo.GetProgress(ctx, req.UserID, req.CourseID)
	if err != nil {
		return res, err
	}

	return res, nil
}
