package abstractfactory

import (
	"fmt"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()

	SetName(string)
	GetName() string
}

type BasePizza struct {
	Name      string
	Dough     *Dough
	Sauce     *Sauce
	Cheese    *Cheese
	Veggies   []*Veggie
	Pepperoni *Pepperoni
	Clam      *Clam
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

func (b *BasePizza) SetName(name string) {
	b.Name = name
}

func (b *BasePizza) GetName() string {
	return b.Name
}

type CheesePizza struct {
	BasePizza
	IngredientFactory PizzaIngredientFactory
}

func NewCheesePizza(ingredentFactory PizzaIngredientFactory) Pizza {
	return &CheesePizza{
		BasePizza:         BasePizza{Name: "Cheese Pizza"},
		IngredientFactory: ingredentFactory,
	}
}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	c.Dough = c.IngredientFactory.CreateDough()
	c.Sauce = c.IngredientFactory.CreateSauce()
	c.Cheese = c.IngredientFactory.CreateCheese()
}

type ClamPizza struct {
	BasePizza
	IngredientFactory PizzaIngredientFactory
}

func NewClamPizza(ingredentFactory PizzaIngredientFactory) Pizza {
	return &ClamPizza{
		BasePizza:         BasePizza{Name: "Clam Pizza"},
		IngredientFactory: ingredentFactory,
	}
}

func (c *ClamPizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	c.Dough = c.IngredientFactory.CreateDough()
	c.Sauce = c.IngredientFactory.CreateSauce()
	c.Cheese = c.IngredientFactory.CreateCheese()
	c.Clam = c.IngredientFactory.CreateClam()
}
