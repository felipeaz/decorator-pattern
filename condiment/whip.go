package condiment

import "decorator-pattern/beverage"

type Whip struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewWhip(b beverage.Beverage, p float64) *Whip {
	return &Whip{
		Beverage:    b,
		Price:       p,
		Description: "Milk",
	}
}

func (w *Whip) GetPrice() {
	//TODO implement me
	panic("implement me")
}

func (w *Whip) GetDescription() {
	//TODO implement me
	panic("implement me")
}
