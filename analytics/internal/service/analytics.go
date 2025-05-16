package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/glamostoffer/arete/analytics/internal/service/dto"
)

type service struct {
	cfg  Config
	repo repository
}

func New(cfg Config, repo repository) *service {
	return &service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *service) ProcessEvent(ctx context.Context) error {
	event, err := s.repo.SelectAndLockEvent(ctx, s.cfg.MaxAttempts, s.cfg.EventLockTime.Duration)
	if err != nil {
		return err
	}

	switch event.Key {
	case domain.QuizComplete:
		return s.processQuizzEvent(ctx, event)
	case domain.TaskComplete:
		return s.processTaskEvent(ctx, event)
	default:
		return errors.New("UNDEFINED_EVENT")
	}
}

func (s *service) processQuizzEvent(ctx context.Context, event domain.Event) error {
	var payload domain.QuizzEventData
	err := json.Unmarshal(event.Payload, &payload)
	if err != nil {
		return err
	}

	stats, err := s.repo.GetUserStats(ctx, payload.UserID)
	if err != nil {
		return err
	}

	err = s.recalculateUserStats(ctx, event.Key, event.Payload, &stats)
	if err != nil {
		return err
	}

	err = s.repo.UpdateUserStats(ctx, stats)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) processTaskEvent(ctx context.Context, event domain.Event) error {
	return nil
}

func (s *service) recalculateUserStats(ctx context.Context, eventType string, eventPayload json.RawMessage, stats *domain.UserStats) error {
	return nil
}

func (s *service) GetUserStats(ctx context.Context, req dto.GetUserStatsRequest) (res dto.GetUserStatsResponse, err error) {
	stats, err := s.repo.GetUserStats(ctx, req.UserID)
	if err != nil {
		return res, err
	}

	res.UserStats = stats
	return res, nil
}
