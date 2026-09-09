package dto

type CreatePageView struct {
	SessionID       string `json:"session_id" db:"session_id"`
	VisitorHash     string `json:"visitor_hash" db:"visitor_hash"`
	PagePath        string `json:"page_path" db:"page_path"`
	Referrer        string `json:"referrer" db:"referrer"`
	UtmSource       string `json:"utm_source" db:"utm_source"`
	UtmMedium       string `json:"utm_medium" db:"utm_medium"`
	UtmCampaign     string `json:"utm_campaign" db:"utm_campaign"`
	UtmTerm         string `json:"utm_term" db:"utm_term"`
	UtmContent      string `json:"utm_content" db:"utm_content"`
	DeviceType      string `json:"device_type" db:"device_type"`
	Browser         string `json:"browser" db:"browser"`
	OS              string `json:"os" db:"os"`
	DurationSeconds int    `json:"duration_seconds" db:"duration_seconds"`
}

type ListPageViewsParams struct {
	Cursor int64 `json:"cursor" query:"cursor" db:"cursor"`
	Limit  int   `json:"limit" query:"limit" db:"limit"`
}
