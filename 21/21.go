package main

import "fmt"

// Адаптируемый метод, для которого будем строить адаптер
type Adaptee struct{}

// Интерфейс, под который хотим подогнать нашу структуру
type Target interface {
	Operation()
}

// Реализация интерфеса, приватным полем которой является адаптируемая структура
type ConcreteAdapter struct {
	adaptee *Adaptee // private поле
}

// Гарантия на этапе компиляции
var _ Target = (*ConcreteAdapter)(nil)

// Метод адаптер, которым подменяем существующий
func (ca *ConcreteAdapter) Operation() {
	fmt.Println("Adapter method called")
	if ca.adaptee == nil {
		fmt.Println("adaptee is nil; skipping")
		return
	}
	ca.adaptee.MyOperation()
}

// Реальный метод
func (a *Adaptee) MyOperation() {
	fmt.Printf("Called MyOperation method, of Adaptee struct\n")
}

// Функция, реализующая подмену существующего класса, на целевой интерфейс
// Путем создания структуры - адаптера
func NewAdapter(a *Adaptee) Target {
	return &ConcreteAdapter{a}
}

func main() {
	a := NewAdapter(&Adaptee{})

	a.Operation()

	// Данный паттерн применяется, в построении гексогональной архитектуры,
	// Когда для доступа между слоями существуют порты и адаптеры,
	// Например, для сохранения обратной совместимости,
	// 			 или независимости, между модулями
	// 			 или отделения бизнес логики от инфры

	// Из плюсов можно выделить:
	// 1) Изоляцию преобразования сущностей в адаптере
	// 2) Не нежно переписывать ВСЮ логику под новую структуру
	// 3) Упрощение тестирования, проще встроить моки etc

	// Из минусов:
	// 1) Сложность между адаптаций новой структуры в старую логику
	// 2) ?

	// Примеры:
	// 1) http.HandleFunc
	// 2) bufio.Reader
	// 3) часто скрывают реализацию логгера
}
