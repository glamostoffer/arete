package eventprocessor

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

func (p *processor) handleQuizEvent(data json.RawMessage, userID int64, courseID uuid.UUID) error {
	return nil
}

func (p *processor) handleTaskEvent(data json.RawMessage, userID int64, courseID uuid.UUID) error {
	return nil
}
