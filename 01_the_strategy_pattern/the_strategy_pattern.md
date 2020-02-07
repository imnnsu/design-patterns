### It started with a simple SimUDuck app

```go
type Duck struct {}

func (d *Duck) Quack() {
    fmt.Println("Quack!")
}

func (d *Duck) Swim() {
    fmt.Println("Swim!")
}

func (d *Duck) Display() {
    fmt.Println("Duck!")
}

type MallardDuck struct {
    Duck
}

func (m *MallardDuck) Display() {
    fmt.Println("MallardDuck!")
}

type RedheadDuck struct {
    Duck
}

func (r *RedheadDuck) Display() {
    fmt.Println("RedheadDuck!")
}
```

### But now we need the ducks to FLY

```go
type Duck struct {}

func (d *Duck) Swim() {
    fmt.Println("Swim!")
}

func (d *Duck) Display() {
    fmt.Println("Duck!")
}

type Flyale interface {
    Fly()
}

type Quackable interface {
    Quack()
}

type MallardDuck struct {
    Duck
}

func (m *MallardDuck) Display() {
    fmt.Println("MallardDuck!")
}

func (m *MallardDuck) Fly() {
    fmt.Println("Fly!")
}

func (m *MallardDuck) Quack() {
    fmt.Println("Quack!")
}

type RedheadDuck struct {
    Duck
}

func (r *RedheadDuck) Display() {
    fmt.Println("RedheadDuck!")
}

func (m RedheadDuck) Fly() {
    fmt.Println("Fly!")
}

func (m RedheadDuck) Quack() {
    fmt.Println("Quack!")
}

type RubberDuck struct {
    Duck
}

func (r *RubberDuck) Display() {
    fmt.Println("RubberDuck!")
}

func (r *RubberDuck) Quack() {
    fmt.Println("Quack!")
}

type DecoyDuck struct {
    Duck
}

func (r *DecoyDuck) Display() {
    fmt.Println("DecoyDuck!")
}
```

> We know that not all of the subclasses should have ﬂying or quacking behavior, so inheritance isn’t the right answer. But while having the subclasses implement Flyable and/or Quackable solves part of the problem (no inappropriately ﬂying rubber ducks), it completely destroys code reuse for those behaviors, so it just creates a different maintenance nightmare. And of course there might be more than one kind of ﬂying behavior even among the ducks that do ﬂy...

### Zeroing in on the problem...

> **Design Principle** Identify the aspects of your application that vary and separate them from what stays the same.

### Designing the Duck Behaviors

> **Design Principle** Program to an interface, not an implementation.

### Integrating the Duck Behavior

```go
type FlyBehavior interface {
    Fly()
}

type FlyWithWings struct {}

func (f *FlyWithWings) Fly() {
    fmt.Println("Fly with wings!")
}

type FlyNoWay struct {}

func (f *FlyNoWay) Fly() {}

type QuackBehavior interface{
    Quack()
}

type Quack struct {}

func (q *Quack) Quack() {
    fmt.Println("Quack!")
}

type Squeak struct {}

func (s *Squeak) Quack() {
    fmt.Println("Squeak!")
}

type MuteQuack struct {}

func (m *MuteQuack) Quack() {}

type Duck struct {
    FlyBehavior FlyBehavior
    QuackBehavior QuackBehavior
}

func (d *Duck) Swim() {
    fmt.Println("Swim!")
}

func (d *Duck) Display() {
    fmt.Println("Duck!")
}

func (d *Duck) PerformQuack() {
    d.QuackBehavior.Quack()
}

func (d *Duck) SetQuackBehavior(quackBehavior QuackBehavior) {
    d.QuackBehavior = quackBehavior
}

func (d *Duck) PerformFly() {
    d.FlyBehavior.Fly()
}

func (d *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
    d.FlyBehavior = flyBehavior
}

type MallardDuck struct {
    Duck
}

func NewMallardDuck() MallardDuck {
    return MallardDuck {
        Duck: Duck {
            FlyBehavior: FlyWithWings{},
            QuackBehavior: Quack{},
        },
    }
}

func (m *MallardDuck) Display() {
    fmt.Println("MallardDuck!")
}
```

### HAS-A can be better than IS-A

> **Desgin Principle** Favor composition over inheritance.

### Speaking of Design Patterns...

> **The Strategy Pattern** defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.
