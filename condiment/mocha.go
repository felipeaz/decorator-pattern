package condiment

import "decorator-pattern/beverage"

type Mocha struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewMocha(b beverage.Beverage, p float64) *Mocha {
	return &Mocha{
		Beverage:    b,
		Price:       p,
		Description: "Milk",
	}
}

func (m *Mocha) GetPrice() {
	//TODO implement me
	panic("implement me")
}

func (m *Mocha) GetDescription() {
	//TODO implement me
	panic("implement me")
}
