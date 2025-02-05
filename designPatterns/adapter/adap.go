// адаптер

package main

import "fmt"

type Target interface {
	Operation()
}

type Adaptee struct {
}

func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am the best")
}

type ConcreteAdapter struct {
	*Adaptee
}

func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

func NewAdapter(adaptee *Adaptee) Target {
	return &ConcreteAdapter{adaptee}
}

// основной метод для демонстрации
func main() {
	fmt.Println("Adapter demo:")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
