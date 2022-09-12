package condiment

import (
	"decorator-pattern/beverage"
	"fmt"
)

type Mocha struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewMocha(b beverage.Beverage) beverage.Beverage {
	return &Mocha{
		Beverage:    b,
		Price:       MochaPrice,
		Description: "Mocha",
	}
}

func (m *Mocha) GetPrice() float64 {
	return m.Price + m.Beverage.GetPrice()
}

func (m *Mocha) GetDescription() string {
	return fmt.Sprintf("%s %s", m.Beverage.GetDescription(), m.Description)
}
