// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h *Human) SayName() {
	fmt.Printf("My name is %s.\n", h.Name)
}

func (h *Human) SayAge() {
	fmt.Printf("I am %d years old.\n", h.Age)
}

func (h *Human) SayGender() {
	fmt.Printf("I'm a %s.\n", h.Gender)
}

type Action struct {
	Human
}

func (act *Action) WelcomeAnyone(name string) {
	fmt.Printf("Hello, %s!\n", name)
	act.SayName()
	act.SayAge()
	act.SayGender()
}

func main() {
	person := Human{
		Name:   "Nikita",
		Age:    26,
		Gender: "male",
	}
	var act = Action{person}
	act.WelcomeAnyone("Anton")
}
