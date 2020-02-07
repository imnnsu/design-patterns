package abstractfactory

type PizzaIngredientFactory interface {
	CreateDough() *Dough
	CreateSauce() *Sauce
	CreateCheese() *Cheese
	CreateVeggies() []*Veggie
	CreatePepperoni() *Pepperoni
	CreateClam() *Clam
}

type NYPizzaIngredientFactory struct{}

func NewNYPizzaIngredientFactory() *NYPizzaIngredientFactory {
	return &NYPizzaIngredientFactory{}
}

func (n *NYPizzaIngredientFactory) CreateDough() *Dough {
	return NewThinCrustDough()
}

func (n *NYPizzaIngredientFactory) CreateSauce() *Sauce {
	return NewMarinaraSauce()
}

func (n *NYPizzaIngredientFactory) CreateCheese() *Cheese {
	return NewReggianoCheese()
}

func (n *NYPizzaIngredientFactory) CreateVeggies() []*Veggie {
	return []*Veggie{NewGarlic(), NewOnion(), NewMushroom(), NewRedPepper()}
}

func (n *NYPizzaIngredientFactory) CreatePepperoni() *Pepperoni {
	return NewSlicedPepperoni()
}

func (n *NYPizzaIngredientFactory) CreateClam() *Clam {
	return NewFreshClams()
}
