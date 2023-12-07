package service

import (
	"httpservermodule/calendar/entity"
	"httpservermodule/calendar/repository"
)

type EventService struct {
	repo *repository.EventRepo
}

func NewEventService(repo *repository.EventRepo) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(event entity.Event) entity.Event {
	return s.repo.Add(event)
}

func (s *EventService) DeleteById(id int) {
	s.repo.DeleteById(id)
}

func (s *EventService) Update(event entity.Event) entity.Event {
	return s.repo.Update(event)
}

func (s *EventService) ExistsById(id int) bool {
	return s.repo.ExistsById(id)
}

func (s *EventService) GetForDay() []entity.Event {
	return s.repo.GetForDay()
}

func (s *EventService) GetForWeek() []entity.Event {
	return s.repo.GetForWeek()
}

func (s *EventService) GetForMonth() []entity.Event {
	return s.repo.GetForMonth()
}
