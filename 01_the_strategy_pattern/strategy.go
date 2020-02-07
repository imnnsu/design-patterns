package main

import (
	"fmt"
)

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("Fly with wings!")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {}

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack!")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Squeak!")
}

type MuteQuack struct{}

func (m *MuteQuack) Quack() {}

type Duck struct {
	FlyBehavior   FlyBehavior
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
	return MallardDuck{
		Duck: Duck{
			FlyBehavior:   &FlyWithWings{},
			QuackBehavior: &Quack{},
		},
	}
}

func (m *MallardDuck) Display() {
	fmt.Println("MallardDuck!")
}

func main() {
	duck := NewMallardDuck()
	duck.Display()
	duck.PerformFly()
	duck.PerformQuack()
}
