package dto

type NameCountPair struct {
	Name  string `json:"name" db:"name"`
	Count int64  `json:"count" db:"count"`
}

type DailyMetric struct {
	Date      string `json:"date" db:"date"`
	PageViews int64  `json:"page_views" db:"page_views"`
	Leads     int64  `json:"leads" db:"leads"`
}

type MetricsSummary struct {
	TotalPageViews int64           `json:"total_page_views"`
	UniqueVisitors int64           `json:"unique_visitors"`
	TotalLeads     int64           `json:"total_leads"`
	TotalEvents    int64           `json:"total_events"`
	ConversionRate float64         `json:"conversion_rate"`
	TopUtmSources  []NameCountPair `json:"top_utm_sources"`
	TopDevices     []NameCountPair `json:"top_devices"`
	TopEvents      []NameCountPair `json:"top_events"`
	TopLeadSources []NameCountPair `json:"top_lead_sources"`
	DailyTrends    []DailyMetric   `json:"daily_trends"`
}

