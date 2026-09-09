package api

import (
	"net/http"

	"go-landing-page/internal/repositories"

	"github.com/labstack/echo"
)

type MetricsAPI struct {
	metricsRepository *repositories.MetricsRepository
}

func NewMetricsAPI(metricsRepository *repositories.MetricsRepository) *MetricsAPI {
	return &MetricsAPI{metricsRepository: metricsRepository}
}

func (m *MetricsAPI) GetSummary(c echo.Context) error {
	summary, err := m.metricsRepository.GetSummary()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve metrics summary")
	}

	return c.JSON(http.StatusOK, summary)
}
