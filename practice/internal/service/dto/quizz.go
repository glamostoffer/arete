package dto

import (
	"github.com/glamostoffer/arete/practice/internal/domain"
	v1 "github.com/glamostoffer/arete/practice/pkg/api/grpc/v1"
	"github.com/gofrs/uuid"
)

type GetCourseQuizzesRequest struct {
	CourseID int64
	UserID   int64
	Limit    int64
	Offset   int64
}
type GetCourseQuizzesResponse struct {
	Quizzes []domain.Quizz
}

func (r *GetCourseQuizzesResponse) ToProto() *v1.GetCourseQuizzesResponse {
	pbQuizzes := make([]*v1.Quizz, 0, len(r.Quizzes))

	for _, quizz := range r.Quizzes {
		pbQuizzes = append(pbQuizzes, &v1.Quizz{
			Id:          quizz.ID,
			CourseID:    quizz.CourseID,
			Title:       quizz.Title,
			Description: quizz.Description,
			IsFinished:  quizz.IsFinished,
			IsLocked:    quizz.IsLocked,
		})
	}

	return &v1.GetCourseQuizzesResponse{
		Quizzes: pbQuizzes,
	}
}

type StartQuizzRequest struct {
	UserID   int64
	QuizzID  int64
	CourseID int64
}
type StartQuizzResponse struct {
	SessionID      uuid.UUID
	TotalQuestions int64
	FirstQuestion  domain.QuestionWithOpts
}

func (r *StartQuizzResponse) ToProto() *v1.StartQuizzResponse {
	pbOptions := make([]*v1.QuestionOption, 0, len(r.FirstQuestion.Options))

	for _, opt := range r.FirstQuestion.Options {
		pbOptions = append(pbOptions, &v1.QuestionOption{
			Id:   opt.ID,
			Text: opt.OptionText,
		})
	}

	pbQuestion := &v1.Question{
		Id:      r.FirstQuestion.Question.ID,
		QuizzID: r.FirstQuestion.Question.QuizzID,
		Text:    r.FirstQuestion.Question.QuestionText,
		Options: pbOptions,
	}

	return &v1.StartQuizzResponse{
		SessionID:      r.SessionID.String(),
		TotalQuestions: r.TotalQuestions,
		FirstQuestion:  pbQuestion,
	}
}

type SubmitQuizzQuestionRequest struct {
	QuestionID int64
	OptionID   int64
	SessionID  uuid.UUID
}
type SubmitQuizzQuestionResponse struct {
	IsCorrect   bool
	Explanation string
}

func (r *SubmitQuizzQuestionResponse) ToProto() *v1.SubmitQuizzQuestionResponse {
	return &v1.SubmitQuizzQuestionResponse{
		IsCorrect:   r.IsCorrect,
		Explanation: r.Explanation,
	}
}

type GetQuizzQuestionRequest struct {
	SessionID uuid.UUID
}
type GetQuizzQuestionResponse struct {
	Question domain.QuestionWithOpts
}

func (r *GetQuizzQuestionResponse) ToProto() *v1.GetQuizzQuestionResponse {
	pbOptions := make([]*v1.QuestionOption, 0, len(r.Question.Options))

	for _, opt := range r.Question.Options {
		pbOptions = append(pbOptions, &v1.QuestionOption{
			Id:   opt.ID,
			Text: opt.OptionText,
		})
	}

	pbQuestion := &v1.Question{
		Id:      r.Question.Question.ID,
		QuizzID: r.Question.Question.QuizzID,
		Text:    r.Question.Question.QuestionText,
		Options: pbOptions,
	}

	return &v1.GetQuizzQuestionResponse{
		Question: pbQuestion,
	}
}

type SubmitQuizzRequest struct {
	SessionID uuid.UUID
}
type SubmitQuizzResponse struct {
	RightAnswersCount int64
	IsFinished        bool
}

func (r *SubmitQuizzResponse) ToProto() *v1.SubmitQuizzResponse {
	return &v1.SubmitQuizzResponse{
		RightAnswers: r.RightAnswersCount,
		IsFinished:   r.IsFinished,
	}
}
