package condiment

import "decorator-pattern/beverage"

type Soy struct {
	Beverage    beverage.Beverage
	Price       float64
	Description string
}

func NewSoy(b beverage.Beverage, p float64) *Soy {
	return &Soy{
		Beverage:    b,
		Price:       p,
		Description: "Milk",
	}
}

func (s *Soy) GetPrice() {
	//TODO implement me
	panic("implement me")
}

func (s *Soy) GetDescription() {
	//TODO implement me
	panic("implement me")
}
