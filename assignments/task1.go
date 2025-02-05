package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Println("Hello, I'm a human, my name is", h.Name)
}

type Action struct {
	Human
	ActionName string
}

func (a Action) fulfill() {
	fmt.Println(a.ActionName, "I'm doing something")
}

func main() {
	a := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
		ActionName: "Run",
	}

	a.Greet()
	a.fulfill()
}
