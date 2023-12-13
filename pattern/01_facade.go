package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Паттерн Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку. Он упрощает взаимодействие клиента с системой, предоставляя ему уровень абстракции.

Применимость паттерна:
	Когда необходимо предоставить простой интерфейс для сложной системы.
	Когда нужно уменьшить зависимости между клиентом и системой компонентов.

Плюсы паттерна "Фасад":
	Упрощение использования системы: Фасад предоставляет упрощенный интерфейс, что упрощает работу с системой.
	Сокрытие сложности: Клиент взаимодействует только с фасадом, не заботясь о внутренней структуре системы.
	Уменьшение зависимостей: Фасад позволяет клиенту взаимодействовать с системой, не зная всех ее деталей.

Минусы паттерна "Фасад":
	Ограниченность гибкости: Если клиенту нужен более сложный или специфичный функционал, он может столкнуться с ограничениями упрощенного интерфейса фасада.
	Дублирование функционала: Иногда возможно дублирование функционала между фасадом и конкретными компонентами системы, что может привести к несогласованности.

	Пример:
*/

// EmailService - сервис для отправки email-уведомлений
type EmailService struct{}

// SMSService - сервис для отправки SMS-уведомлений
type SMSService struct{}

// PushNotificationService - сервис для отправки push-уведомлений
type PushNotificationService struct{}

// NotificationFacade - фасад для управления уведомлениями
type NotificationFacade struct {
	emailService            *EmailService
	smsService              *SMSService
	pushNotificationService *PushNotificationService
}

func NewNotificationFacade() *NotificationFacade {
	return &NotificationFacade{
		emailService:            &EmailService{},
		smsService:              &SMSService{},
		pushNotificationService: &PushNotificationService{},
	}
}

// SendNotification - метод фасада для отправки уведомления клиенту
func (facade *NotificationFacade) SendNotification(message string) {
	// Отправка уведомления через email
	facade.emailService.sendEmail(message)

	// Отправка уведомления через SMS
	facade.smsService.sendSMS(message)

	// Отправка уведомления через push-уведомления
	facade.pushNotificationService.sendPushNotification(message)
}

func (es *EmailService) sendEmail(message string) {
	fmt.Println("Email:", message)
}

func (ss *SMSService) sendSMS(message string) {
	fmt.Println("SMS:", message)
}

func (pns *PushNotificationService) sendPushNotification(message string) {
	fmt.Println("Push notification:", message)
}

func main() {
	// Использование фасада для отправки уведомлений
	notificationFacade := NewNotificationFacade()

	message := "Важное сообщение!"
	notificationFacade.SendNotification(message)
}
