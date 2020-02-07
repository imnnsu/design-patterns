package abstractfactory

type Clam struct {
	Name string
}

type Dough struct {
	Name string
}

type Sauce struct {
	Name string
}

type Cheese struct {
	Name string
}

type Pepperoni struct {
	Name string
}

type Veggie struct {
	Name string
}

func NewThinCrustDough() *Dough {
	return &Dough{Name: "Thin Crust Dough"}
}

func NewMarinaraSauce() *Sauce {
	return &Sauce{Name: "Marinara Sauce"}
}

func NewReggianoCheese() *Cheese {
	return &Cheese{Name: "Reggiano Cheese"}
}

func NewGarlic() *Veggie {
	return &Veggie{Name: "Garlic"}
}
func NewOnion() *Veggie {
	return &Veggie{Name: "Onion"}
}
func NewMushroom() *Veggie {
	return &Veggie{Name: "Mushroom"}
}
func NewRedPepper() *Veggie {
	return &Veggie{Name: "Pepper"}
}

func NewSlicedPepperoni() *Pepperoni {
	return &Pepperoni{Name: "Sliced Pepperoni"}
}

func NewFreshClams() *Clam {
	return &Clam{Name: "Fresh Clams"}
}
