package models

import "time"

type Lead struct {
	ID          int64     `db:"id" json:"id"`
	SessionID   string    `db:"session_id" json:"session_id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	Company     string    `db:"company" json:"company"`
	Phone       string    `db:"phone" json:"phone"`
	Description string    `db:"description" json:"description"`
	UtmSource   string    `db:"utm_source" json:"utm_source"`
	UtmCampaign string    `db:"utm_campaign" json:"utm_campaign"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
