package domain

type Quizz struct {
	ID             int64
	CourseID       int64
	Title          string
	Description    string
	PassingScore   int64 // кол-во необходимых правильных ответов для прохождения
	SequenceNumber int64 // порядковый номер квиза в рамках курса
	IsLocked       bool
	IsFinished     bool
}

type Question struct {
	ID           int64
	QuizzID      int64
	QuestionText string
	Explanation  string
}

type QuestionOption struct {
	ID         int64
	QuestionID int64
	OptionText string
	IsCorrect  bool
}

type QuestionWithOpts struct {
	Question Question
	Options  []QuestionOption
}

type QuizzSessionResult struct {
	QuestionsTotal    int64
	RightAnswersCount int64
}

type QuizzFinishedEvent struct {
	UserID   int64
	QuizzID  int64
	CourseID int64
}
