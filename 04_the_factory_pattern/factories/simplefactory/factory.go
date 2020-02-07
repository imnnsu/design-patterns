package simplefactory

import (
	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/pizza"
)

type SimplePizzaFactory struct{}

func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

func (f *SimplePizzaFactory) CreatePizza(pizzaType string) pizza.Pizza {
	switch pizzaType {
	case "cheese":
		return pizza.NewCheesePizza()
	case "pepperoni":
		return pizza.NewPepperoniPizza()
	case "clam":
		return pizza.NewClamPizza()
	case "veggie":
		return pizza.NewVeggiePizza()
	default:
		return pizza.NewCheesePizza()
	}
}
