package handlers

import (
	"encoding/json"
	"errors"
	"httpservermodule/calendar/dto"
	"httpservermodule/calendar/mapper"
	"httpservermodule/calendar/service"
	"io"
	"log"
	"net/http"
)

type CalendarEventHandler struct {
	mapper  *mapper.EventMapper
	service *service.EventService
}

func NewCalendarEventHandler(service *service.EventService, mapper *mapper.EventMapper) *CalendarEventHandler {
	return &CalendarEventHandler{service: service, mapper: mapper}
}

// CreateEventHandler обрабатывает запросы создания события
func (h *CalendarEventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	defer h.recoverServerError(w)
	body := readRequestBody(r.Body)
	var eventDto dto.CreateEvent
	if err := json.Unmarshal(body, &eventDto); err != nil {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	event, err := h.mapper.CreateToEntity(eventDto)
	if err != nil {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	event = h.service.Create(event)
	h.respondWithResult(w, 200, event)
}

func (h *CalendarEventHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	defer h.recoverServerError(w)
	body := readRequestBody(r.Body)
	var eventDto dto.UpdateEvent
	if err := json.Unmarshal(body, &eventDto); err != nil {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	if !h.service.ExistsById(eventDto.ID) {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	event, err := h.mapper.UpdateToEntity(eventDto)
	if err != nil {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	event = h.service.Update(event)
	h.respondWithResult(w, 200, event)
}

func (h *CalendarEventHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	defer h.recoverServerError(w)
	body := readRequestBody(r.Body)
	var eventDto dto.UpdateEvent
	if err := json.Unmarshal(body, &eventDto); err != nil {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	if !h.service.ExistsById(eventDto.ID) {
		h.respondWithError(w, 400, errors.New("bad request"))
		return
	}
	h.service.DeleteById(eventDto.ID)
	h.respondWithResult(w, 200, "")
}

func (h *CalendarEventHandler) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	events := h.service.GetForDay()
	h.respondWithResult(w, 200, events)
}

func (h *CalendarEventHandler) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	events := h.service.GetForWeek()
	h.respondWithResult(w, 200, events)
}

func (h *CalendarEventHandler) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	events := h.service.GetForMonth()
	h.respondWithResult(w, 200, events)
}

// readRequestBody считывает тело запроса
func readRequestBody(reqBody io.ReadCloser) []byte {
	defer func(reqBody io.ReadCloser) {
		if err := reqBody.Close(); err != nil {
			log.Panic(err)
		}
	}(reqBody)
	body := make([]byte, 100)
	for {
		body = body[:cap(body)]
		n, err := reqBody.Read(body)
		body = body[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
	}
	return body
}

func (h *CalendarEventHandler) recoverServerError(w http.ResponseWriter) {
	if rec := recover(); rec != nil {
		h.respondWithError(w, 500, errors.New("internal server error"))
		return
	}
}

func (h *CalendarEventHandler) respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		h.respondWithError(w, 500, errors.New("internal server error"))
		return
	}
}

func (h *CalendarEventHandler) respondWithResult(w http.ResponseWriter, status int, data interface{}) {
	h.respondWithJSON(w, status, map[string]interface{}{"result": data})
}

func (h *CalendarEventHandler) respondWithError(w http.ResponseWriter, status int, err error) {
	h.respondWithJSON(w, status, map[string]string{"error": err.Error()})
}
