package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpservermodule/calendar/mapper"
	"httpservermodule/calendar/repository"
	"httpservermodule/calendar/service"
	"httpservermodule/handlers"
	"httpservermodule/middleware"
	"log"
	"net/http"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	r := mux.NewRouter()

	eventRepo := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepo)
	eventMapper := mapper.NewEventMapper()
	calendarEventHandler := handlers.NewCalendarEventHandler(eventService, eventMapper)

	// Обработчики HTTP-методов
	r.HandleFunc("/create_event", calendarEventHandler.CreateEventHandler).Methods("POST")
	r.HandleFunc("/update_event", calendarEventHandler.UpdateEventHandler).Methods("POST")
	r.HandleFunc("/delete_event", calendarEventHandler.DeleteEventHandler).Methods("POST")
	r.HandleFunc("/events_for_day", calendarEventHandler.EventsForDayHandler).Methods("GET")
	r.HandleFunc("/events_for_week", calendarEventHandler.EventsForWeekHandler).Methods("GET")
	r.HandleFunc("/events_for_month", calendarEventHandler.EventsForMonthHandler).Methods("GET")

	// Применение middleware для логирования
	r.Use(middleware.LoggingMiddleware)

	port := 8080
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting server on :%d...\n", port)
	log.Fatal(server.ListenAndServe())
}
