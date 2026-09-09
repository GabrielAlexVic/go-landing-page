package dto

type CreateLead struct {
	SessionID   string `json:"session_id" db:"session_id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Company     string `json:"company" db:"company"`
	Phone       string `json:"phone" db:"phone"`
	Description string `json:"description" db:"description"`
	UtmSource   string `json:"utm_source" db:"utm_source"`
	UtmCampaign string `json:"utm_campaign" db:"utm_campaign"`
}

type ListLeadsParams struct {
	Cursor int64 `json:"cursor" query:"cursor" db:"cursor"`
	Limit  int   `json:"limit" query:"limit" db:"limit"`
}
