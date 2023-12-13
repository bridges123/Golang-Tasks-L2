package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн "Команда" — это паттерн проектирования, который превращает запросы или простые операции в объекты.
Позволяет параметризовать клиентов с различными запросами, организовывать очередь запросов и поддерживать отмену операций.

Применимость паттерна:
	Когда вы хотите параметризовать объекты с операциями.
	Когда вы хотите передавать операции как параметры.
	Когда вы хотите поддерживать отмену операций и выполнение их в различные моменты времени.

Плюсы паттерна "Команда":
	Отделение отправителя от получателя: Команды обеспечивают отделение объектов, отправляющих запросы, от объектов, которые их обрабатывают.
	Поддержка отмены и повтора операций: Команды могут поддерживать отмену операций, а также повтор их выполнения.
	Легкость добавления новых команд: Добавление новых команд не требует изменения существующего кода отправителей.

Минусы паттерна "Команда":
	Увеличение числа классов: Использование паттерна может привести к созданию множества классов, особенно если существует большое количество различных команд.
	Усложнение кода: Введение команд может усложнить код и сделать его менее читаемым.

Пример:
*/

// Command - интерфейс команды
type Command interface {
	Execute()
}

type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is off")
}

// LightOnCommand - конкретная команда для включения светильника
type LightOnCommand struct {
	Light *Light
}

func (lc *LightOnCommand) Execute() {
	lc.Light.TurnOn()
}

// LightOffCommand - конкретная команда для выключения светильника
type LightOffCommand struct {
	Light *Light
}

func (lc *LightOffCommand) Execute() {
	lc.Light.TurnOff()
}

type Door struct{}

func (d *Door) Open() {
	fmt.Println("Door is open")
}

func (d *Door) Close() {
	fmt.Println("Door is closed")
}

// DoorOpenCommand - конкретная команда для открытия двери
type DoorOpenCommand struct {
	Door *Door
}

func (doc *DoorOpenCommand) Execute() {
	doc.Door.Open()
}

// DoorCloseCommand - конкретная команда для закрытия двери
type DoorCloseCommand struct {
	Door *Door
}

func (dcc *DoorCloseCommand) Execute() {
	dcc.Door.Close()
}

// RemoteControl - пульт управления
type RemoteControl struct {
	Command Command
}

// PressButton - нажатие кнопки на пульте
func (rc *RemoteControl) PressButton() {
	rc.Command.Execute()
}

func main() {
	light := &Light{}
	lightOnCommand := &LightOnCommand{Light: light}
	lightOffCommand := &LightOffCommand{Light: light}

	door := &Door{}
	doorOpenCommand := &DoorOpenCommand{Door: door}
	doorCloseCommand := &DoorCloseCommand{Door: door}

	remote := &RemoteControl{}

	// Привязка команд к пульту
	remote.Command = lightOnCommand
	remote.PressButton()

	remote.Command = lightOffCommand
	remote.PressButton()

	remote.Command = doorOpenCommand
	remote.PressButton()

	remote.Command = doorCloseCommand
	remote.PressButton()
}
