package mapper

import (
	"errors"
	"httpservermodule/calendar/dto"
	"httpservermodule/calendar/entity"
	"time"
)

type EventMapper struct {
}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) CreateToEntity(dto dto.CreateEvent) (entity.Event, error) {
	var err error
	date, err := time.Parse(time.DateOnly, dto.Date)
	if err != nil {
		return entity.Event{}, errors.New("invalid data")
	}

	event := entity.Event{UserID: dto.UserID, Title: dto.Title, Date: date}
	return event, nil
}

func (m *EventMapper) UpdateToEntity(dto dto.UpdateEvent) (entity.Event, error) {
	var err error
	date, err := time.Parse(time.DateOnly, dto.Date)
	if err != nil {
		return entity.Event{}, errors.New("invalid data")
	}

	event := entity.Event{ID: dto.ID, UserID: dto.UserID, Title: dto.Title, Date: date}
	return event, nil
}
