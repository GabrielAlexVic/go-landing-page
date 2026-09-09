package repositories

import (
	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
	"go-landing-page/internal/models"
)

type LeadRepository struct {
	db database.Service
}

func NewLeadRepository(db database.Service) *LeadRepository {
	return &LeadRepository{db: db}
}

func (r *LeadRepository) CreateLead(lead dto.CreateLead) error {
	_, err := r.db.GetDB().NamedExec(queries.CreateLead, lead)
	if err != nil {
		return err
	}

	return nil
}

func (r *LeadRepository) ListLeads(params dto.ListLeadsParams) ([]models.Lead, error) {
	var leads []models.Lead

	limit := params.Limit
	if limit <= 0 {
		limit = 50
	}

	err := r.db.GetDB().Select(&leads, queries.ListLeads, params.Cursor, limit)
	if err != nil {
		return nil, err
	}
	return leads, nil
}
