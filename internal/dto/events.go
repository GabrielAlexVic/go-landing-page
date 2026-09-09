package dto

import "encoding/json"

type CreateEvent struct {
	SessionID  string          `json:"session_id" db:"session_id"`
	EventType  string          `json:"event_type" db:"event_type"`
	EventLabel string          `json:"event_label" db:"event_label"`
	Metadata   json.RawMessage `json:"metadata" db:"metadata"`
}

type ListEventsParams struct {
	Cursor int64 `json:"cursor" query:"cursor" db:"cursor"`
	Limit  int   `json:"limit" query:"limit" db:"limit"`
}
