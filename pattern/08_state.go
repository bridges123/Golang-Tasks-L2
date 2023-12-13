package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Паттерн "Состояние" — это паттерн проектирования, который позволяет объекту изменять свое поведение при изменении его внутреннего состояния. Превращает каждое состояние объекта в отдельный класс.

Применимость паттерна:
	Когда поведение объекта зависит от его состояния и может изменяться во время выполнения.
	Когда у вас есть множество условных операторов, зависящих от состояния объекта.

Плюсы паттерна "Состояние":
	Изолирование состояний: Каждое состояние инкапсулируется в отдельном классе, что облегчает добавление новых состояний и изменение поведения.
	Упрощение кода контекста: Исключает длинные цепочки условий в контексте, связанные с различными состояниями.
	Гибкость: Позволяет легко добавлять новые состояния и изменять поведение объекта, не изменяя его код.

Минусы паттерна "Состояние":
	Увеличение числа классов: Введение отдельных классов для каждого состояния может привести к увеличению числа классов в системе.
	Сложность понимания: Некоторые аспекты системы могут стать сложными для понимания из-за децентрализации логики состояний.

Пример:
*/

// State - интерфейс для состояний телевизора
type State interface {
	PowerOn()
	PowerOff()
	ChannelUp()
	ChannelDown()
}

type TV struct {
	state State
}

// SetState - установка текущего состояния телевизора
func (tv *TV) SetState(state State) {
	tv.state = state
}

func (tv *TV) PowerOn() {
	tv.state.PowerOn()
}

func (tv *TV) PowerOff() {
	tv.state.PowerOff()
}

func (tv *TV) ChannelUp() {
	tv.state.ChannelUp()
}

func (tv *TV) ChannelDown() {
	tv.state.ChannelDown()
}

type OffState struct{}

func (os *OffState) PowerOn() {
	fmt.Println("Turning the TV on.")
}

func (os *OffState) PowerOff() {
	fmt.Println("The TV is already off.")
}

func (os *OffState) ChannelUp() {
	fmt.Println("Cannot change channel, the TV is off.")
}

func (os *OffState) ChannelDown() {
	fmt.Println("Cannot change channel, the TV is off.")
}

type OnState struct{}

func (os *OnState) PowerOn() {
	fmt.Println("The TV is already on.")
}

func (os *OnState) PowerOff() {
	fmt.Println("Turning the TV off.")
}

func (os *OnState) ChannelUp() {
	fmt.Println("Changing channel up.")
}

func (os *OnState) ChannelDown() {
	fmt.Println("Changing channel down.")
}

func main() {
	tv := &TV{}
	tv.SetState(&OffState{})

	// Включение и переключение канала
	tv.PowerOn()
	tv.ChannelUp()

	// Выключение и повторное включение
	tv.PowerOff()
	tv.PowerOn()

	// Переключение канала после включения
	tv.ChannelDown()
	tv.ChannelUp()
}
