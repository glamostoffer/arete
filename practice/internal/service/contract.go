package service

import (
	"context"
	"time"

	"github.com/glamostoffer/arete/practice/internal/domain"
	"github.com/gofrs/uuid"
)

type repository interface {
	SelectQuizzesByCourseID(ctx context.Context, courseID, userID, limit, offset int64) (res []domain.Quizz, err error)
	GetQuizz(ctx context.Context, quizzID, userID int64) (domain.Quizz, error)
	SelectQuizzQuestions(ctx context.Context, quizzID int64) (res []domain.Question, err error)
	SelectQuestionOptions(ctx context.Context, questionID int64) (res []domain.QuestionOption, err error)
	SelectAllQuizzOptions(ctx context.Context, questionIDs []int64) (res []domain.QuestionOption, err error)
	SelectQuestionsWithOptions(ctx context.Context, quizzID int64) (map[domain.Question][]domain.QuestionOption, error)
	MarkQuizzCompleted(ctx context.Context, userID, quizzID int64) error
}

type cache interface {
	SetQuizzSession(
		ctx context.Context,
		info domain.QuizzFinishedEvent,
		sessionID uuid.UUID,
		questionWithOpts map[domain.Question][]domain.QuestionOption,
		ttl time.Duration,
	) error
	GetQuestionFromSession(
		ctx context.Context,
		sessionID uuid.UUID,
		questionID int64,
	) (res domain.QuestionWithOpts, err error)
	GetDelQuestionFromSession(
		ctx context.Context,
		sessionID uuid.UUID,
		questionID int64,
	) (res domain.QuestionWithOpts, err error)
	AddRightAnswerToQuizzSessionResult(
		ctx context.Context,
		sessionID uuid.UUID,
	) error
	GetRandomQuestionFromSession(
		ctx context.Context,
		sessionID uuid.UUID,
	) (res domain.QuestionWithOpts, err error)
	GetInfoFromSession(
		ctx context.Context,
		sessionID uuid.UUID,
	) (res domain.QuizzSessionResult, info domain.QuizzFinishedEvent, err error)
}

type producer interface {
	WriteMessage(ctx context.Context, key, value string) error
}
