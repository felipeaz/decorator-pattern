package condiment

import "decorator-pattern/beverage"

type Milk struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewMilk(b beverage.Beverage, p float64) *Milk {
	return &Milk{
		Beverage:    b,
		Price:       p,
		Description: "Milk",
	}
}

func (c *Milk) GetPrice() {
	//TODO implement me
	panic("implement me")
}

func (c *Milk) GetDescription() {
	//TODO implement me
	panic("implement me")
}
