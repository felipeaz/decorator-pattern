package condiment

import (
	"decorator-pattern/beverage"
	"fmt"
)

type Soy struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewSoy(b beverage.Beverage) beverage.Beverage {
	return &Soy{
		Beverage:    b,
		Price:       SoyPrice,
		Description: "Soy",
	}
}

func (s *Soy) GetPrice() float64 {
	return s.Price + s.Beverage.GetPrice()
}

func (s *Soy) GetDescription() string {
	return fmt.Sprintf("%s %s", s.Beverage.GetDescription(), s.Description)
}
