package dto

import "github.com/glamostoffer/arete/practice/internal/domain"

type Quizz struct {
	ID             int64  `db:"id"`
	CourseID       int64  `db:"course_id"`
	Title          string `db:"title"`
	Description    string `db:"description"`
	PassingScore   int64  `db:"passing_score"`
	SequenceNumber int64  `db:"sequence_number"`
	IsLocked       bool   `db:"is_locked"`
	IsFinished     bool   `db:"is_finished"`
}

func QuizzFromDomain(in domain.Quizz) Quizz {
	return Quizz{
		ID:             in.ID,
		CourseID:       in.CourseID,
		Title:          in.Title,
		Description:    in.Description,
		PassingScore:   in.PassingScore,
		SequenceNumber: in.SequenceNumber,
		IsLocked:       in.IsLocked,
		IsFinished:     in.IsFinished,
	}
}

func (q Quizz) ToDomain() domain.Quizz {
	return domain.Quizz{
		ID:             q.ID,
		CourseID:       q.CourseID,
		Title:          q.Title,
		Description:    q.Description,
		PassingScore:   q.PassingScore,
		SequenceNumber: q.SequenceNumber,
		IsLocked:       q.IsLocked,
		IsFinished:     q.IsFinished,
	}
}

type Question struct {
	ID           int64  `db:"id"`
	QuizzID      int64  `db:"quizz_id"`
	QuestionText string `db:"question"`
	Explanation  string `db:"explanation"`
}

func QuestionFromDomain(in domain.Question) Question {
	return Question{
		ID:           in.ID,
		QuizzID:      in.QuizzID,
		QuestionText: in.QuestionText,
		Explanation:  in.Explanation,
	}
}

func (q Question) ToDomain() domain.Question {
	return domain.Question{
		ID:           q.ID,
		QuizzID:      q.QuizzID,
		QuestionText: q.QuestionText,
		Explanation:  q.Explanation,
	}
}

type QuestionOption struct {
	ID         int64  `db:"id"`
	QuestionID int64  `db:"question_id"`
	OptionText string `db:"option"`
	IsCorrect  bool   `db:"is_correct"`
}

func QuestionOptionFromDomain(in domain.QuestionOption) QuestionOption {
	return QuestionOption{
		ID:         in.ID,
		QuestionID: in.QuestionID,
		OptionText: in.OptionText,
		IsCorrect:  in.IsCorrect,
	}
}

func (qo QuestionOption) ToDomain() domain.QuestionOption {
	return domain.QuestionOption{
		ID:         qo.ID,
		QuestionID: qo.QuestionID,
		OptionText: qo.OptionText,
		IsCorrect:  qo.IsCorrect,
	}
}
