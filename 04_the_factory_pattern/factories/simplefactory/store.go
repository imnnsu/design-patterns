package simplefactory

import (
	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/pizza"
)

type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{factory: NewSimplePizzaFactory()}
}

func (p *PizzaStore) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := p.factory.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
