package abstractfactory

type PizzaStore interface {
	OrderPizza(self PizzaStore, pizzaType string) Pizza
	CreatePizza(pizzaType string) Pizza
}

type BasePizzaStore struct{}

func (b *BasePizzaStore) OrderPizza(self PizzaStore, pizzaType string) Pizza {
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

func (n *NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	ingredientFactory := NewNYPizzaIngredientFactory()
	var pizza Pizza

	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("New York Style Cheese Pizza")
	case "clam":
		pizza = NewClamPizza(ingredientFactory)
		pizza.SetName("New York Style Clam Pizza")
	//case "veggie":
	//case "pepperoni":
	}

	return pizza
}
