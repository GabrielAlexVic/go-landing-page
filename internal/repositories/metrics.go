package repositories

import (
	"math"

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

	return summary, nil
}
