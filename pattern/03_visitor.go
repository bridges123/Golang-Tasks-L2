package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн "Посетитель" — это паттерн проектирования, который позволяет определить новую операцию без изменения классов объектов, к которым она применяется.
Этот паттерн полезен, когда у вас есть набор объектов, и вы хотите выполнить над ними операции, но не хотите добавлять новые методы в их интерфейсы.

Применимость паттерна:
	Когда у вас есть множество объектов с разными типами, и вы хотите выполнить операции над ними, не изменяя их код.
	Когда операции, выполняемые над объектами, часто меняются или расширяются.

Плюсы паттерна "Посетитель":
	Отделение операций от структуры объекта: Позволяет добавлять новые операции, не изменяя структуры объектов.
	Поддержка операций для разных типов объектов: Посетитель позволяет добавлять операции для различных типов объектов, не изменяя сами объекты.

Минусы паттерна "Посетитель":
	Усложнение кода: Внедрение паттерна может привести к усложнению кода из-за создания множества классов, связанных с посетителем.
	Нарушение инкапсуляции: Посетитель требует открытия методов объектов для посетителя, что может нарушить инкапсуляцию.

Пример:
*/

type Shape interface {
	Accept(Visitor)
}

type Circle struct {
	Radius float64
}

// Accept - принятие посетителя для круга
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Square struct {
	SideLength float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

// Visitor - интерфейс для посетителя
type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

// AreaCalculator - посетитель для вычисления площади фигур
type AreaCalculator struct {
	TotalArea float64
}

// VisitCircle - вычисление площади круга
func (ac *AreaCalculator) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Calculating area for Circle: %f\n", area)
	ac.TotalArea += area
}

// VisitSquare - вычисление площади квадрата
func (ac *AreaCalculator) VisitSquare(s *Square) {
	area := s.SideLength * s.SideLength
	fmt.Printf("Calculating area for Square: %f\n", area)
	ac.TotalArea += area
}

func main() {
	circle := &Circle{Radius: 5.0}
	square := &Square{SideLength: 4.0}

	areaCalculator := &AreaCalculator{}

	// Применение посетителя для вычисления площади фигур
	circle.Accept(areaCalculator)
	square.Accept(areaCalculator)

	fmt.Printf("Total Area: %f\n", areaCalculator.TotalArea)
}
