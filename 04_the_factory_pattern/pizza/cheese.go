package pizza

type CheesePizza struct {
	*BasePizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{
		BasePizza: NewBasePizza("Cheese Pizza", []string{"Cheese"}),
	}
}

type NYCheesePizza struct {
	*BasePizza
}

func NewNYCheesePizza() *NYCheesePizza {
	return &NYCheesePizza{
		BasePizza: NewBasePizza("NY Style Sauce and Cheese Pizza",
			[]string{"Grated Reggiano Cheese"}),
	}
}

type ChicagoStylePizza struct {
	*BasePizza
}

func NewChicagoStylePizza() *ChicagoStylePizza {
	return &ChicagoStylePizza{
		BasePizza: NewBasePizza("Chicago Style Deep Dish Cheese Pizza",
			[]string{"Shredded Mozzarella Cheese"}),
	}
}
