package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Паттерн "Фабричный метод" — это паттерн проектирования, который определяет интерфейс для создания объекта,
но оставляет подклассам решение о том, какой класс инстанцировать.
Таким образом, он делегирует ответственность по созданию объектов подклассам.

Применимость паттерна:
	Когда у вас есть некоторый общий интерфейс для создания объектов, но конкретные классы решают, какой объект создавать.
	Когда система должна быть независимой от процесса создания и композиции объектов.
	Когда создание объекта — это детали реализации и они должны быть скрыты от клиентского кода.

Плюсы паттерна "Фабричный метод":
	Отделение создания объекта от использования: Клиентский код использует интерфейс, а не конкретные классы, что упрощает изменения в системе.
	Расширяемость: Позволяет добавлять новые подклассы, не изменяя существующий код.
	Контроль над созданием объектов: Каждый подкласс может решать, какой именно объект ему создавать.

Минусы паттерна "Фабричный метод":
	Усложнение структуры: Добавление новых классов может привести к усложнению структуры системы.
	Невозможность гарантировать уникальность создаваемых объектов: Клиент может получить разные объекты при использовании различных подклассов, что может быть нежелательным.

Пример:
*/

type Logger interface {
	Log(message string)
}

// FileLogger - конкретный логгер для записи в файл
type FileLogger struct{}

func (fl *FileLogger) Log(message string) {
	fmt.Println("Log to file:", message)
}

// ConsoleLogger - конкретный логгер для вывода в консоль
type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(message string) {
	fmt.Println("Log to console:", message)
}

// LoggerFactory - интерфейс для создания логгера
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory - фабрика для создания файлового логгера
type FileLoggerFactory struct{}

func (flf *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory - фабрика для создания консольного логгера
type ConsoleLoggerFactory struct{}

func (clf *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

func main() {
	fileLoggerFactory := &FileLoggerFactory{}
	consoleLoggerFactory := &ConsoleLoggerFactory{}

	// Создание файлового логгера
	fileLogger := fileLoggerFactory.CreateLogger()
	fileLogger.Log("This message goes to a file")

	// Создание консольного логгера
	consoleLogger := consoleLoggerFactory.CreateLogger()
	consoleLogger.Log("This message goes to the console")
}
