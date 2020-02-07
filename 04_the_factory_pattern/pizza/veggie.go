package pizza

type VeggiePizza struct {
	*BasePizza
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		BasePizza: NewBasePizza("Veggie Pizza", []string{"Veggie"}),
	}
}

type NYVeggiePizza struct {
	*BasePizza
}

func NewNYVeggiePizza() *NYVeggiePizza {
	return &NYVeggiePizza{
		BasePizza: NewBasePizza("NY Veggie Pizza",
			[]string{"NY Veggie"}),
	}
}
