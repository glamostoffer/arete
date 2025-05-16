package service

import (
	"context"
	"encoding/json"

	"github.com/glamostoffer/arete/practice/internal/domain"
	"github.com/glamostoffer/arete/practice/internal/service/dto"
	"github.com/gofrs/uuid"
)

func (s *service) GetCourseQuizzes(
	ctx context.Context,
	req dto.GetCourseQuizzesRequest,
) (res dto.GetCourseQuizzesResponse, err error) {
	quizzes, err := s.repo.SelectQuizzesByCourseID(
		ctx,
		req.CourseID,
		req.UserID,
		req.Limit,
		req.Offset,
	)
	if err != nil {
		return res, err
	}

	res.Quizzes = quizzes
	return res, nil
}

func (s *service) StartQuizz(
	ctx context.Context,
	req dto.StartQuizzRequest,
) (res dto.StartQuizzResponse, err error) {
	questsWithOpts, err := s.repo.SelectQuestionsWithOptions(ctx, req.QuizzID)
	if err != nil {
		return res, err
	}

	sessionID, err := uuid.NewV7()
	if err != nil {
		return res, err
	}

	err = s.cache.SetQuizzSession(
		ctx,
		domain.QuizzFinishedEvent{
			UserID:   0,
			QuizzID:  0,
			CourseID: 0,
		},
		sessionID,
		questsWithOpts,
		s.cfg.QuizzSessionTTL.Duration,
	)
	if err != nil {
		return res, err
	}

	for q, o := range questsWithOpts {
		res.FirstQuestion.Question = q
		res.FirstQuestion.Options = o
		break
	}
	res.SessionID = sessionID
	res.TotalQuestions = int64(len(questsWithOpts))

	return res, nil
}

func (s *service) SubmitQuizzQuestion(
	ctx context.Context,
	req dto.SubmitQuizzQuestionRequest,
) (res dto.SubmitQuizzQuestionResponse, err error) {
	questionWithOpts, err := s.cache.GetDelQuestionFromSession(ctx, req.SessionID, req.QuestionID)
	if err != nil {
		return res, err
	}

	for _, opt := range questionWithOpts.Options {
		if req.OptionID == opt.ID {
			res.IsCorrect = opt.IsCorrect
			if opt.IsCorrect {
				err = s.cache.AddRightAnswerToQuizzSessionResult(ctx, req.SessionID)
				if err != nil {
					return res, err
				}
			}
		}
	}

	res.Explanation = questionWithOpts.Question.Explanation

	return res, nil
}

func (s *service) GetQuizzQuestion(
	ctx context.Context,
	req dto.GetQuizzQuestionRequest,
) (res dto.GetQuizzQuestionResponse, err error) {
	questionWithOpts, err := s.cache.GetRandomQuestionFromSession(ctx, req.SessionID)
	if err != nil {
		return res, err
	}

	res.Question = questionWithOpts
	return res, nil
}

func (s *service) SubmitQuizz(
	ctx context.Context,
	req dto.SubmitQuizzRequest,
) (res dto.SubmitQuizzResponse, err error) {
	result, info, err := s.cache.GetInfoFromSession(ctx, req.SessionID)
	if err != nil {
		return res, err
	}

	quizz, err := s.repo.GetQuizz(ctx, info.QuizzID, info.UserID)
	if err != nil {
		return res, err
	}

	res.RightAnswersCount = result.RightAnswersCount
	if result.RightAnswersCount >= quizz.PassingScore {
		res.IsFinished = true
		err = s.repo.MarkQuizzCompleted(ctx, info.QuizzID, info.UserID)
		if err != nil {
			return res, err
		}

		data, err := json.Marshal(info)
		if err != nil {
			return res, err
		}

		s.producer.WriteMessage(ctx, "quizz-finished", string(data))
	}

	return res, nil
}
