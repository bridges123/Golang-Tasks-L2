package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Паттерн "Стратегия" — это паттерн проектирования, который определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми.
Позволяет изменять алгоритмы независимо от клиентов, которые их используют.

Применимость паттерна:
	Когда есть несколько вариантов реализации для какой-то задачи, и они могут быть выбраны во время выполнения.
	Когда нужно предоставить возможность клиентам выбирать алгоритм из набора доступных.

Плюсы паттерна "Стратегия":
	Изменение поведения на лету: Позволяет изменять алгоритмы на лету, не затрагивая код клиента.
	Избегание дублирования кода: Разделяет различные варианты реализации, избегая дублирования кода.
	Уменьшение зависимостей: Позволяет отделить алгоритмы от кода, который их использует, уменьшая зависимости.

Минусы паттерна "Стратегия":
	Усложнение кода: Может привести к увеличению числа классов в системе из-за введения отдельных стратегий.
	Возможность забыть установить стратегию: Клиент должен явным образом выбрать стратегию, и забыть сделать это может привести к неправильному поведению.

Пример:
*/

// PaymentStrategy - интерфейс стратегии оплаты
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardPayment - конкретная стратегия оплаты кредитной картой
type CreditCardPayment struct{}

func (ccp *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card\n", amount)
}

// PayPalPayment - конкретная стратегия оплаты через PayPal
type PayPalPayment struct{}

func (ppp *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal\n", amount)
}

type Order struct {
	amount          float64
	paymentStrategy PaymentStrategy
}

// ProcessOrder - обработка заказа с использованием выбранной стратегии оплаты
func (o *Order) ProcessOrder() {
	o.paymentStrategy.Pay(o.amount)
}

func main() {
	order1 := &Order{amount: 150.0, paymentStrategy: &CreditCardPayment{}}
	order2 := &Order{amount: 100.0, paymentStrategy: &PayPalPayment{}}

	order1.ProcessOrder()
	order2.ProcessOrder()
}
