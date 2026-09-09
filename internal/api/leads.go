package api

import (
	"go-landing-page/internal/dto"
	"go-landing-page/internal/repositories"
	"net/http"

	"github.com/labstack/echo"
)

type LeadApi struct {
	leadeRepository *repositories.LeadRepository
}

func NewLeadApi(leadeRepository *repositories.LeadRepository) *LeadApi {
	return &LeadApi{leadeRepository: leadeRepository}
}

func (l *LeadApi) CreateLead(c echo.Context) error {
	body := new(dto.CreateLead)
	if err := c.Bind(body); err != nil {
		return err
	}

	err := l.leadeRepository.CreateLead(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create lead")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "lead created successfully"})
}

func (l *LeadApi) ListLeads(c echo.Context) error {
	body := new(dto.ListLeadsParams)
	if err := c.Bind(body); err != nil {
		return err
	}

	leads, err := l.leadeRepository.ListLeads(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to list leads: "+err.Error())
	}

	return c.JSON(http.StatusOK, leads)
}
