package grpc

import (
	"context"

	"github.com/glamostoffer/arete/practice/internal/service/dto"
	v1 "github.com/glamostoffer/arete/practice/pkg/api/grpc/v1"
	"github.com/gofrs/uuid"
)

type handler struct {
	v1.UnimplementedPracticeServer

	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetCourseQuizzes(
	ctx context.Context,
	req *v1.GetCourseQuizzesRequest,
) (res *v1.GetCourseQuizzesResponse, err error) {
	out, err := h.service.GetCourseQuizzes(ctx, dto.GetCourseQuizzesRequest{
		CourseID: req.GetCourseID(),
		UserID:   req.GetUserID(),
		Limit:    req.GetLimit(),
		Offset:   req.GetOffset(),
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) StartQuizz(
	ctx context.Context,
	req *v1.StartQuizzRequest,
) (res *v1.StartQuizzResponse, err error) {
	out, err := h.service.StartQuizz(ctx, dto.StartQuizzRequest{
		UserID:   req.GetUserID(),
		QuizzID:  req.GetQuizzID(),
		CourseID: req.GetCourseID(),
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) SubmitQuizzQuestion(
	ctx context.Context,
	req *v1.SubmitQuizzQuestionRequest,
) (res *v1.SubmitQuizzQuestionResponse, err error) {
	sessionID, err := uuid.FromString(req.GetSessionID())
	if err != nil {
		return res, err
	}

	out, err := h.service.SubmitQuizzQuestion(ctx, dto.SubmitQuizzQuestionRequest{
		QuestionID: req.GetQuestionID(),
		OptionID:   req.GetOptionID(),
		SessionID:  sessionID,
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) GetQuizzQuestion(
	ctx context.Context,
	req *v1.GetQuizzQuestionRequest,
) (res *v1.GetQuizzQuestionResponse, err error) {
	sessionID, err := uuid.FromString(req.GetSessionID())
	if err != nil {
		return res, err
	}

	out, err := h.service.GetQuizzQuestion(ctx, dto.GetQuizzQuestionRequest{
		SessionID: sessionID,
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) SubmitQuizz(
	ctx context.Context,
	req *v1.SubmitQuizzRequest,
) (res *v1.SubmitQuizzResponse, err error) {
	sessionID, err := uuid.FromString(req.GetSessionID())
	if err != nil {
		return res, err
	}

	out, err := h.service.SubmitQuizz(ctx, dto.SubmitQuizzRequest{
		SessionID: sessionID,
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}
