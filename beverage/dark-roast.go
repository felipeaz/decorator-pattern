package beverage

type DarkRoast struct {
	Price       float64
	Description string
}

func NewDarkRoast() Beverage {
	return &DarkRoast{
		Price:       DarkRoastPrice,
		Description: "Dark Roast",
	}
}

func (d *DarkRoast) GetPrice() float64 {
	return d.Price
}

func (d *DarkRoast) GetDescription() string {
	return d.Description
}
