package domain

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type EventType string

const (
	QuizComplete = EventType("quiz_complete")
	TaskComplete = EventType("task_complete")
)

type Rating struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      int64     `json:"userID" db:"user_id"`
	CourseID    uuid.UUID `json:"courseID" db:"course_id"`
	Score       int64     `json:"score" db:"score"`
	LastUpdated time.Time `json:"lastUpdated" db:"last_updated"`
}

type Progress struct {
	ID                      uuid.UUID       `json:"id" db:"id"`
	UserID                  int64           `json:"userID" db:"user_id"`
	CourseID                uuid.UUID       `json:"courseID" db:"course_id"`
	TotalMaterialsCompleted int64           `json:"totalMaterialsCompleted" db:"total_matetials_completed"`
	TotalTasksCompleted     int64           `json:"totalTasksCompleted" db:"total_tasks_completed"`
	TotalQuizzesCompleted   int64           `json:"totalQuizzesCompletes" db:"total_quizzes_completed"`
	TotalScore              int64           `json:"totalScore" db:"total_score"`
	CompletionRate          decimal.Decimal `json:"completionRate" db:"completion_rate"`
	LastUpdated             time.Time       `json:"lastUpdated" db:"last_updated"`
}

type Event struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	UserID      int64           `json:"userID" db:"user_id"`
	CourseID    uuid.UUID       `json:"courseID" db:"course_id"`
	Type        EventType       `json:"type" db:"type"`
	Data        json.RawMessage `json:"data" db:"data"`
	CreatedAt   time.Time       `json:"createdAt" db:"created_at"`
	ProcessedAt *time.Time      `json:"-" db:"processed_at"`
}

func (e *Event) ParseKafkaMessage(*sarama.ConsumerMessage) error {
	return nil
}
