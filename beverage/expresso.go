package beverage

type Expresso struct {
	Price       float64
	Description string
}

func NewExpresso() Beverage {
	return &Expresso{
		Price:       ExpressoPrice,
		Description: "Expresso",
	}
}

func (e *Expresso) GetPrice() float64 {
	return e.Price
}

func (e *Expresso) GetDescription() string {
	return e.Description
}
