package repositories

import (
	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
	"go-landing-page/internal/models"
)

type PageViewRepository struct {
	db database.Service
}

func NewPageViewRepository(db database.Service) *PageViewRepository {
	return &PageViewRepository{db: db}
}

func (r *PageViewRepository) CreatePageView(pageView dto.CreatePageView) error {
	_, err := r.db.GetDB().NamedExec(queries.CreatePageView, pageView)
	if err != nil {
		return err
	}

	return nil
}

func (r *PageViewRepository) ListPageViews(params dto.ListPageViewsParams) ([]models.PageView, error) {
	var pageViews []models.PageView

	limit := params.Limit
	if limit <= 0 {
		limit = 50
	}

	err := r.db.GetDB().Select(&pageViews, queries.ListPageViews, params.Cursor, limit)
	if err != nil {
		return nil, err
	}
	return pageViews, nil
}
