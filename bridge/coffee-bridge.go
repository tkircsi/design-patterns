package bridge

import "fmt"

func MainCoffeeBridge() {
	largeCoffeeWithMilk := LargeCoffee{
		additives: &Milk{},
	}
	largeCoffeeWithMilk.Order(2)

	smallCoffeWithSugar := SmallCoffee{
		additives: &Sugar{},
	}
	smallCoffeWithSugar.Order(1)

}

type Coffee interface {
	Order(count int)
}

type Additives interface {
	AddAdditives()
}

type LargeCoffee struct {
	additives Additives
}

func (c *LargeCoffee) Order(count int) {
	fmt.Print("Large coffee ")
	c.additives.AddAdditives()
	fmt.Println(count, "cup")
}

type SmallCoffee struct {
	additives Additives
}

func (c *SmallCoffee) Order(count int) {
	fmt.Print("Small coffe ")
	c.additives.AddAdditives()
	fmt.Println(count, "cup")
}

type Milk struct{}

func (m *Milk) AddAdditives() {
	fmt.Print("With Milk ")
}

type Sugar struct{}

func (s *Sugar) AddAdditives() {
	fmt.Print("With Sugar ")
}
