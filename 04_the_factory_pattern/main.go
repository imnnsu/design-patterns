package main

import (
	"fmt"

	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/factories/abstractfactory"
	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/factories/factorymethod"
	"github.com/minsunchina/go-design-patterns/04_the_factory_pattern/factories/simplefactory"
)

func main() {
	// Simple factory
	simpleFactoryStore := simplefactory.NewPizzaStore()
	pizza := simpleFactoryStore.OrderPizza("cheese")
	fmt.Println("Steve ordered a " + pizza.GetName())
	fmt.Println("---")

	// Factory method
	factoryMethodStore := factorymethod.NewNYPizzaStore()
	pizza = factoryMethodStore.OrderPizza(factoryMethodStore, "cheese")
	fmt.Println("Ethan ordered a " + pizza.GetName())
	fmt.Println("---")

	// Abstract factory
	abstractFactoryStore := abstractfactory.NewNYPizzaStore()
	pizza = abstractFactoryStore.OrderPizza(abstractFactoryStore, "cheese")
	fmt.Println("Joel ordered a " + pizza.GetName())
}
