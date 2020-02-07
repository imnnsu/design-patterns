package pizza

type PepperoniPizza struct {
	*BasePizza
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{
		BasePizza: NewBasePizza("Pepperoni Pizza", []string{"Pepperoni"}),
	}
}

type NYPepperoniPizza struct {
	*BasePizza
}

func NewNYPepperoniPizza() *NYPepperoniPizza {
	return &NYPepperoniPizza{
		BasePizza: NewBasePizza("NY Pepperoni Pizza",
			[]string{"NY Pepperoni"}),
	}
}
