package pizza

import (
	"fmt"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()

	GetName() string
}

type BasePizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func NewBasePizza(name string, toppings []string) *BasePizza {
	return &BasePizza{
		Name:     name,
		Toppings: toppings,
	}
}

func (b *BasePizza) Prepare() {
	fmt.Println("Preparing " + b.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings: ", b.Toppings)
}

func (b *BasePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (b *BasePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (b *BasePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (b *BasePizza) GetName() string {
	return b.Name
}
