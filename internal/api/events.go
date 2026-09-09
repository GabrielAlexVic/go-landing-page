package api

import (
	"net/http"

	"go-landing-page/internal/dto"
	"go-landing-page/internal/repositories"

	"github.com/labstack/echo"
)

type EventAPI struct {
	eventRepository *repositories.EventRepository
}

func NewEventAPI(eventRepository *repositories.EventRepository) *EventAPI {
	return &EventAPI{eventRepository: eventRepository}
}

func (e *EventAPI) CreateEvent(c echo.Context) error {
	body := new(dto.CreateEvent)
	if err := c.Bind(body); err != nil {
		return err
	}

	err := e.eventRepository.CreateEvent(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create event: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "event created successfully"})
}

func (e *EventAPI) ListEvents(c echo.Context) error {
	body := new(dto.ListEventsParams)
	if err := c.Bind(body); err != nil {
		return err
	}

	events, err := e.eventRepository.ListEvents(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to list events")
	}

	return c.JSON(http.StatusOK, events)
}
