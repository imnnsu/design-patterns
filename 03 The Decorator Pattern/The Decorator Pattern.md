### The Open-closed Principle

> **Design Principle** Classes should be open for extension, but closed for modification.

### The Decortor Pattern defined

> **The Decortor Pattern** attaches additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending funtionality.

### Writing the Starbuzz code

```go
type Beverage interface {
    GetDescription() string
    Cost() float64
}

type BaseBeverage struct {
    description string
}

func (b *BaseBeverage) GetDescription() string {
    return b.description
}
```

### Coding beverages

```go
type Espresso struct {
    BaseBeverage
}

func NewEspresso() *Espresso {
    return &Espresso{BaseBeverage{"Espresso"}}
}

func (e *Espresso) Cost() float64 {
    return 1.99
}

type HouseBlend struct {
	BaseBeverage
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{BaseBeverage{"HouseBlend"}}
}

func (h *HouseBlend) Cost() float64 {
	return 0.89
}
```

### Coding condiments

```go
type Mocha struct {
    beverage Beverage
}

func NewMocha(beverage Beverage) *Mocha {
    return &Mocha{beverage}
}

func (m *Mocha) GetDescription() string {
    return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
    return 0.20 + m.beverage.Cost()
}
```

### Serving some coffees

```go
func main() {
    var beverage1, beverage2, beverage3 Beverage
    beverage1 = NewEspresso()
    fmt.Printf("%s $%.2f\n", beverage1.GetDescription(), beverage1.Cost())
    beverage2 = NewMocha(NewEspresso())
    fmt.Printf("%s $%.2f\n", beverage2.GetDescription(), beverage2.Cost())
    beverage3 = NewMocha(NewMocha(NewHouseBlend()))
    fmt.Printf("%s $%.2f\n", beverage3.GetDescription(), beverage3.Cost())
}
```