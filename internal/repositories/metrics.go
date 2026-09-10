package repositories

import (
	"math"
	"sort"

	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
)

type MetricsRepository struct {
	db database.Service
}

func NewMetricsRepository(db database.Service) *MetricsRepository {
	return &MetricsRepository{db: db}
}

func (r *MetricsRepository) GetSummary() (dto.MetricsSummary, error) {
	var summary dto.MetricsSummary

	db := r.db.GetDB()

	err := db.Get(&summary.TotalPageViews, queries.GetTotalPageViews)
	if err != nil {
		return summary, err
	}

	err = db.Get(&summary.UniqueVisitors, queries.GetUniqueVisitors)
	if err != nil {
		return summary, err
	}

	err = db.Get(&summary.TotalLeads, queries.GetTotalLeads)
	if err != nil {
		return summary, err
	}

	err = db.Get(&summary.TotalEvents, queries.GetTotalEvents)
	if err != nil {
		return summary, err
	}

	if summary.UniqueVisitors > 0 {
		cvr := (float64(summary.TotalLeads) / float64(summary.UniqueVisitors)) * 100
		summary.ConversionRate = math.Round(cvr*100) / 100
	}

	var topSources []dto.NameCountPair
	err = db.Select(&topSources, queries.GetTopUtmSources)
	if err == nil {
		summary.TopUtmSources = topSources
	} else {
		summary.TopUtmSources = []dto.NameCountPair{}
	}

	var topDevices []dto.NameCountPair
	err = db.Select(&topDevices, queries.GetTopDevices)
	if err == nil {
		summary.TopDevices = topDevices
	} else {
		summary.TopDevices = []dto.NameCountPair{}
	}

	var topEvents []dto.NameCountPair
	err = db.Select(&topEvents, queries.GetTopEvents)
	if err == nil {
		summary.TopEvents = topEvents
	} else {
		summary.TopEvents = []dto.NameCountPair{}
	}

	var topLeadSources []dto.NameCountPair
	err = db.Select(&topLeadSources, queries.GetTopLeadSources)
	if err == nil {
		summary.TopLeadSources = topLeadSources
	} else {
		summary.TopLeadSources = []dto.NameCountPair{}
	}

	var dailyViews []dto.NameCountPair
	var dailyLeads []dto.NameCountPair

	_ = db.Select(&dailyViews, queries.GetDailyPageViews)
	_ = db.Select(&dailyLeads, queries.GetDailyLeads)

	dateMap := make(map[string]*dto.DailyMetric)

	for _, v := range dailyViews {
		dateMap[v.Name] = &dto.DailyMetric{
			Date:      v.Name,
			PageViews: v.Count,
			Leads:     0,
		}
	}

	for _, l := range dailyLeads {
		if item, exists := dateMap[l.Name]; exists {
			item.Leads = l.Count
		} else {
			dateMap[l.Name] = &dto.DailyMetric{
				Date:      l.Name,
				PageViews: 0,
				Leads:     l.Count,
			}
		}
	}

	var dates []string
	for d := range dateMap {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	dailyTrends := make([]dto.DailyMetric, 0, len(dates))
	for _, d := range dates {
		dailyTrends = append(dailyTrends, *dateMap[d])
	}
	summary.DailyTrends = dailyTrends

	return summary, nil
}

