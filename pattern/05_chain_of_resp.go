package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Паттерн "Цепочка вызовов" — это паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый обработчик решает, может ли он обработать запрос, и либо обрабатывает его, либо передает следующему обработчику в цепочке.

Применимость паттерна:
	Когда есть несколько объектов, которые могут обработать запрос, и порядок обработки не фиксирован.
	Когда вы хотите, чтобы объекты автоматически передавали запросы друг другу.

Плюсы паттерна "Цепочка вызовов":
	Отделение отправителя запроса от получателя: Каждый обработчик может решить, может ли он обработать запрос, и, если нет, передать его дальше по цепочке.
	Гибкость и расширяемость: Легко добавлять новые обработчики и изменять порядок цепочки без изменения отправителя запроса.

Минусы паттерна "Цепочка вызовов":
	Не гарантирует обработку запроса: Нет гарантии, что запрос будет обработан в конечном итоге, что может быть проблемой, если нет явного завершающего обработчика.
	Возможна неопределенность: Если цепочка слишком длинная или не управляется должным образом, это может привести к неопределенности или ошибкам в выполнении программы.

Пример:
*/

// PaymentHandler - интерфейс обработчика оплаты
type PaymentHandler interface {
	SetNext(PaymentHandler)
	Pay(amount float64)
}

// CreditCardHandler - обработчик оплаты кредитной картой
type CreditCardHandler struct {
	next PaymentHandler
}

func (c *CreditCardHandler) SetNext(handler PaymentHandler) {
	c.next = handler
}

// Pay - выполнение оплаты кредитной картой
func (c *CreditCardHandler) Pay(amount float64) {
	if amount <= 100 {
		fmt.Println("Paid $", amount, " using Credit Card")
	} else if c.next != nil {
		fmt.Println("Amount too high for Credit Card. Passing to next handler.")
		c.next.Pay(amount)
	} else {
		fmt.Println("No suitable payment method found.")
	}
}

// WalletHandler - обработчик оплаты электронным кошельком
type WalletHandler struct {
	next PaymentHandler
}

func (w *WalletHandler) SetNext(handler PaymentHandler) {
	w.next = handler
}

// Pay - выполнение оплаты электронным кошельком
func (w *WalletHandler) Pay(amount float64) {
	if amount <= 50 {
		fmt.Println("Paid $", amount, " using Wallet")
	} else if w.next != nil {
		fmt.Println("Amount too high for Wallet. Passing to next handler.")
		w.next.Pay(amount)
	} else {
		fmt.Println("No suitable payment method found.")
	}
}

func main() {
	// Использование паттерна "Цепочка вызовов"
	creditCardHandler := &CreditCardHandler{}
	walletHandler := &WalletHandler{}

	// Установка следующего обработчика
	creditCardHandler.SetNext(walletHandler)

	// Обработка оплаты
	creditCardHandler.Pay(70.0)
}
