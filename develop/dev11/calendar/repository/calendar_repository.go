package repository

import (
	"httpservermodule/calendar/entity"
	"time"
)

type EventRepo struct {
	idIncrement int
	events      map[int]entity.Event
}

func NewEventRepository() *EventRepo {
	return &EventRepo{events: map[int]entity.Event{}, idIncrement: 1}
}

func (r *EventRepo) Add(event entity.Event) entity.Event {
	event.ID = r.idIncrement
	r.events[r.idIncrement] = event
	r.idIncrement++
	return event
}

func (r *EventRepo) DeleteById(id int) {
	delete(r.events, id)
}

func (r *EventRepo) Update(event entity.Event) entity.Event {
	r.events[event.ID] = event
	return event
}

func (r *EventRepo) GetForDay() []entity.Event {
	return r.getByTime(time.Hour, 24)
}

func (r *EventRepo) GetForWeek() []entity.Event {
	return r.getByTime(time.Hour*24, 7)
}

func (r *EventRepo) GetForMonth() []entity.Event {
	return r.getByTime(time.Hour*24*7, 30)
}

func (r *EventRepo) getByTime(hours time.Duration, border time.Duration) []entity.Event {
	now := time.Now()
	eventsForDay := make([]entity.Event, 0, 4)
	for _, event := range r.events {
		if event.Date.Sub(now)/hours <= border {
			eventsForDay = append(eventsForDay, event)
		}
	}
	return eventsForDay
}

func (r *EventRepo) ExistsById(id int) bool {
	_, ok := r.events[id]
	return ok
}
