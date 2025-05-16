package repository

import (
	"context"
	"errors"

	"github.com/glamostoffer/arete/practice/internal/domain"
	"github.com/glamostoffer/arete/practice/internal/repository/dto"
	"github.com/lib/pq"
)

func (r *repository) SelectQuizzesByCourseID(ctx context.Context, courseID, userID, limit, offset int64) (res []domain.Quizz, err error) {
	quizzes := make([]dto.Quizz, 0)

	err = r.db.SelectContext(
		ctx,
		&quizzes,
		querySelectQuizzesByCourseID,
		courseID,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	for _, quizz := range quizzes {
		res = append(res, quizz.ToDomain())
	}

	return res, nil
}

func (r *repository) GetQuizz(ctx context.Context, quizzID, userID int64) (domain.Quizz, error) {
	var quizz dto.Quizz

	err := r.db.GetContext(
		ctx,
		&quizz,
		queryGetQuizz,
		quizzID,
		userID,
	)
	if err != nil {
		return domain.Quizz{}, err
	}

	return quizz.ToDomain(), nil
}

func (r *repository) SelectQuizzQuestions(ctx context.Context, quizzID int64) (res []domain.Question, err error) {
	questions := make([]dto.Question, 0)

	err = r.db.SelectContext(
		ctx,
		&questions,
		querySelectQuizzQuestions,
		quizzID,
	)
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		res = append(res, question.ToDomain())
	}

	return res, nil
}

func (r *repository) SelectQuestionOptions(ctx context.Context, questionID int64) (res []domain.QuestionOption, err error) {
	options := make([]dto.QuestionOption, 0)

	err = r.db.SelectContext(
		ctx,
		&options,
		querySelectQuestionOptions,
		questionID,
	)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		res = append(res, option.ToDomain())
	}

	return res, nil
}

func (r *repository) SelectAllQuizzOptions(ctx context.Context, questionIDs []int64) (res []domain.QuestionOption, err error) {
	options := make([]dto.QuestionOption, 0)

	err = r.db.SelectContext(
		ctx,
		&options,
		querySelectAllQuizzOptions,
		pq.Array(questionIDs),
	)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		res = append(res, option.ToDomain())
	}

	return res, nil
}

func (r *repository) SelectQuestionsWithOptions(ctx context.Context, quizzID int64) (map[domain.Question][]domain.QuestionOption, error) {
	var questions []dto.Question
	err := r.db.SelectContext(ctx, &questions, querySelectQuizzQuestions, quizzID)
	if err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return make(map[domain.Question][]domain.QuestionOption), nil
	}

	questionIDs := make([]int64, len(questions))
	questionMap := make(map[int64]domain.Question, len(questions))
	for i, q := range questions {
		questionIDs[i] = q.ID
		questionMap[q.ID] = q.ToDomain()
	}

	var options []dto.QuestionOption
	err = r.db.SelectContext(ctx, &options, querySelectAllQuizzOptions, pq.Array(questionIDs))
	if err != nil {
		return nil, err
	}

	res := make(map[domain.Question][]domain.QuestionOption, len(questions))
	for _, opt := range options {
		if question, exists := questionMap[opt.QuestionID]; exists {
			res[question] = append(res[question], opt.ToDomain())
		}
	}

	return res, nil
}

func (r *repository) MarkQuizzCompleted(ctx context.Context, userID, quizzID int64) error {
	res, err := r.db.ExecContext(
		ctx,
		queryMarkQuizzCompleted,
		userID,
		quizzID,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("INVALID_AFFECTED_ROWS_COUNT") // todo use errlist
	}

	return nil

}
