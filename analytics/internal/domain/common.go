package domain

import (
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	QuizComplete = "quiz_complete"
	TaskComplete = "task_complete"
)

type TaskEventData struct {
	// todo
}

type QuizzEventData struct {
	UserID int64
	// todo
}

type UserStats struct {
	UserCourseProgress
	CourseRating
	GlobalRating
}

type UserCourseProgress struct {
	UserID               int64
	CourseID             int64
	CompletionPercentage float64
	LastUpdated          time.Time
	CompletedLessons     int64
	CompletedQuizzes     int64
	CompletedTasks       int64
}

type CourseRating struct {
	UserID      int64
	CourseID    int64
	Rating      float64
	Position    int64
	LastUpdated time.Time
}

type GlobalRating struct {
	UserID      int64
	Rating      float64
	Position    int64
	LastUpdated time.Time
}
type Event struct {
	ID             int64
	Key            string
	Topic          string
	Payload        json.RawMessage
	IdempotencyKey string
	CreatedAt      time.Time
	LockedUntil    *time.Time
	Attempts       int64
	Error          *string
	ProcessedAt    *time.Time
}

func EventFromMessage(msg kafka.Message) Event {
	return Event{
		Key:       string(msg.Key),
		Topic:     msg.Topic,
		Payload:   msg.Value,
		CreatedAt: time.Now(),
	}
}
