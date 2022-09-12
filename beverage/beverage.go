package beverage

const (
	DarkRoastPrice  = 4.00
	ExpressoPrice   = 1.25
	HouseBlendPrice = 2.50
)

type Beverage interface {
	GetPrice() float64
	GetDescription() string
}
