package condiment

import (
	"decorator-pattern/beverage"
	"fmt"
)

type Whip struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewWhip(b beverage.Beverage) beverage.Beverage {
	return &Whip{
		Beverage:    b,
		Price:       WhipPrice,
		Description: "Whip",
	}
}

func (w *Whip) GetPrice() float64 {
	return w.Price + w.Beverage.GetPrice()
}

func (w *Whip) GetDescription() string {
	return fmt.Sprintf("%s %s", w.Beverage.GetDescription(), w.Description)
}
