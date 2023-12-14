// Реализовать паттерн «адаптер» на любом примере

package main

// Структура, не удовлетворяющая новому интерфейсу NewInterface, из которой нам нужен метод SomeMethod
type OldFunctionality struct {
}

func (old *OldFunctionality) SomeMethod() {
	println("I am a necessary method from old struct")
}

// Новый интерфейс нашей программы
type NewInterface interface {
	NewMethod()
}

// Адаптер, который реализует метод SomeMethod из OldFunctionality через интерфейсный метод NewMethod,
// подходящий для нового интерфейса
type Adapter struct {
	*OldFunctionality
}

// Новый метод, удовлетворяющий NewInterface, но реализующий метод SomeMethod
func (a *Adapter) NewMethod() {
	a.SomeMethod()
}

func NewAdapter() *Adapter {
	return &Adapter{new(OldFunctionality)}
}

func main() {
	a := NewAdapter()
	ImplementedType := NewInterface(a)
	ImplementedType.NewMethod()
}
