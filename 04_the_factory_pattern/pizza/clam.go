package pizza

type ClamPizza struct {
	*BasePizza
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{
		BasePizza: NewBasePizza("Clam Pizza", []string{"Clam"}),
	}
}

type NYClamPizza struct {
	*BasePizza
}

func NewNYClamPizza() *NYClamPizza {
	return &NYClamPizza{
		BasePizza: NewBasePizza("NY Clam Pizza",
			[]string{"NY Clam"}),
	}
}
