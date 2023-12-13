package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Строитель — это паттерн проектирования, который позволяет создавать сложные объекты шаг за шагом.
Он отделяет конструирование сложного объекта от его представления,
что позволяет использовать один и тот же процесс конструирования для создания различных представлений объекта.

Применимость паттерна:
	Когда процесс конструирования объекта состоит из множества шагов.
	Когда необходимо создавать различные представления одного и того же объекта.

Плюсы паттерна "Строитель":
	Постепенное построение объекта: Процесс конструирования происходит шаг за шагом, что облегчает создание сложных объектов.
	Гибкость: Паттерн "Строитель" позволяет изменять внутреннее представление продукта, не меняя его внешний код.

Минусы паттерна "Строитель":
	Усложнение кода: Внедрение паттерна может привести к усложнению кода из-за создания дополнительных классов и интерфейсов.
	Необходимость создания дополнительных классов: Каждый тип продукта может потребовать создания своего собственного строителя, что может увеличить количество классов в системе.

Пример:
*/

// CarBuilder - интерфейс строителя автомобиля
type CarBuilder interface {
	BuildEngine()
	BuildWheels()
	BuildNavigation()
	BuildAirConditioner()
	GetCar() *Car
}

// Car - структура автомобиля
type Car struct {
	Engine         string
	Wheels         string
	Navigation     string
	AirConditioner string
}

// SedanBuilder - конкретный строитель для седана
type SedanBuilder struct {
	car *Car
}

func NewSedanBuilder() CarBuilder {
	return &SedanBuilder{car: &Car{}}
}

func (sb *SedanBuilder) BuildEngine() {
	sb.car.Engine = "Sedan Engine"
}

func (sb *SedanBuilder) BuildWheels() {
	sb.car.Wheels = "Sedan Wheels"
}

func (sb *SedanBuilder) BuildNavigation() {
	sb.car.Navigation = "Basic Navigation"
}

func (sb *SedanBuilder) BuildAirConditioner() {
	sb.car.AirConditioner = "Standard Air Conditioner"
}

// GetCar - получение итогового автомобиля
func (sb *SedanBuilder) GetCar() *Car {
	return sb.car
}

// Director - директор, который управляет процессом строительства
type Director struct {
	builder CarBuilder
}

// Construct - метод для пошагового конструирования автомобиля
func (d *Director) Construct() {
	d.builder.BuildEngine()
	d.builder.BuildWheels()
	d.builder.BuildNavigation()
	d.builder.BuildAirConditioner()
}

func main() {
	sedanBuilder := NewSedanBuilder()
	director := &Director{builder: sedanBuilder}

	// Построение седана
	director.Construct()
	sedan := sedanBuilder.GetCar()

	// Вывод информации о седане
	fmt.Printf("Sedan Info:\nEngine: %s\nWheels: %s\nNavigation: %s\nAir Conditioner: %s\n",
		sedan.Engine, sedan.Wheels, sedan.Navigation, sedan.AirConditioner)
}
