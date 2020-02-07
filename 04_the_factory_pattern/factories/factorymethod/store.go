package factorymethod

import (
	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/pizza"
)

type PizzaStore interface {
	OrderPizza(self PizzaStore, pizzaType string) pizza.Pizza
	CreatePizza(pizzaType string) pizza.Pizza
}

type BasePizzaStore struct{}

func (b *BasePizzaStore) OrderPizza(self PizzaStore, pizzaType string) pizza.Pizza {
	pizza := self.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type NYPizzaStore struct {
	BasePizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{}
}

func (n *NYPizzaStore) CreatePizza(pizzaType string) pizza.Pizza {
	switch pizzaType {
	case "cheese":
		return pizza.NewNYCheesePizza()
	case "pepperoni":
		return pizza.NewNYPepperoniPizza()
	case "clam":
		return pizza.NewNYClamPizza()
	case "veggie":
		return pizza.NewNYVeggiePizza()
	default:
		return pizza.NewNYCheesePizza()
	}
}
