package repositories

import (
	"encoding/json"
	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
	"go-landing-page/internal/models"
)

type EventRepository struct {
	db database.Service
}

func NewEventRepository(db database.Service) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(event dto.CreateEvent) error {
	if len(event.Metadata) == 0 {
		event.Metadata = json.RawMessage("{}")
	}

	_, err := r.db.GetDB().NamedExec(queries.CreateEvent, event)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) ListEvents(params dto.ListEventsParams) ([]models.Event, error) {
	var events []models.Event

	limit := params.Limit
	if limit <= 0 {
		limit = 50
	}

	err := r.db.GetDB().Select(&events, queries.ListEvents, params.Cursor, limit)
	if err != nil {
		return nil, err
	}
	return events, nil
}
