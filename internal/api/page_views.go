package api

import (
	"go-landing-page/internal/dto"
	"go-landing-page/internal/repositories"
	"net/http"

	"github.com/labstack/echo"
)

type PageViewApi struct {
	pageViewRepository *repositories.PageViewRepository
}

func NewPageViewApi(pageViewRepository *repositories.PageViewRepository) *PageViewApi {
	return &PageViewApi{pageViewRepository: pageViewRepository}
}

func (p *PageViewApi) CreatePageView(c echo.Context) error {
	body := new(dto.CreatePageView)
	if err := c.Bind(body); err != nil {
		return err
	}

	err := p.pageViewRepository.CreatePageView(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create page view")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "page view created successfully"})
}

func (p *PageViewApi) ListPageViews(c echo.Context) error {
	body := new(dto.ListPageViewsParams)
	if err := c.Bind(body); err != nil {
		return err
	}

	pageViews, err := p.pageViewRepository.ListPageViews(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to list page views")
	}

	return c.JSON(http.StatusOK, pageViews)
}
