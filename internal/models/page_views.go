package models

import "time"

type PageView struct {
	ID              int64     `db:"id" json:"id"`
	SessionID       string    `db:"session_id" json:"session_id"`
	VisitorHash     string    `db:"visitor_hash" json:"visitor_hash"`
	PagePath        string    `db:"page_path" json:"page_path"`
	Referrer        string    `db:"referrer" json:"referrer"`
	UtmSource       string    `db:"utm_source" json:"utm_source"`
	UtmMedium       string    `db:"utm_medium" json:"utm_medium"`
	UtmCampaign     string    `db:"utm_campaign" json:"utm_campaign"`
	UtmTerm         string    `db:"utm_term" json:"utm_term"`
	UtmContent      string    `db:"utm_content" json:"utm_content"`
	DeviceType      string    `db:"device_type" json:"device_type"`
	Browser         string    `db:"browser" json:"browser"`
	OS              string    `db:"os" json:"os"`
	DurationSeconds int       `db:"duration_seconds" json:"duration_seconds"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
