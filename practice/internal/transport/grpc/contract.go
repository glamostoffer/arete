package grpc

import (
	"context"

	"github.com/glamostoffer/arete/practice/internal/service/dto"
)

type service interface {
	GetCourseQuizzes(
		ctx context.Context,
		req dto.GetCourseQuizzesRequest,
	) (res dto.GetCourseQuizzesResponse, err error)
	StartQuizz(
		ctx context.Context,
		req dto.StartQuizzRequest,
	) (res dto.StartQuizzResponse, err error)
	SubmitQuizzQuestion(
		ctx context.Context,
		req dto.SubmitQuizzQuestionRequest,
	) (res dto.SubmitQuizzQuestionResponse, err error)
	GetQuizzQuestion(
		ctx context.Context,
		req dto.GetQuizzQuestionRequest,
	) (res dto.GetQuizzQuestionResponse, err error)
	SubmitQuizz(
		ctx context.Context,
		req dto.SubmitQuizzRequest,
	) (res dto.SubmitQuizzResponse, err error)
}
