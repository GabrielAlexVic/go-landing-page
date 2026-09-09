package models

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID         int64           `db:"id" json:"id"`
	SessionID  string          `db:"session_id" json:"session_id"`
	EventType  string          `db:"event_type" json:"event_type"`
	EventLabel string          `db:"event_label" json:"event_label"`
	Metadata   json.RawMessage `db:"metadata" json:"metadata"`
	CreatedAt  time.Time       `db:"created_at" json:"created_at"`
}
