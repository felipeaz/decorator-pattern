package condiment

import (
	"decorator-pattern/beverage"
	"fmt"
)

type Milk struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewMilk(b beverage.Beverage) beverage.Beverage {
	return &Milk{
		Beverage:    b,
		Price:       MilkPrice,
		Description: "Milk",
	}
}

func (c *Milk) GetPrice() float64 {
	return c.Price + c.Beverage.GetPrice()
}

func (c *Milk) GetDescription() string {
	return fmt.Sprintf("%s %s", c.Beverage.GetDescription(), c.Description)
}
