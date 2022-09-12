package beverage

type HouseBlend struct {
	Price       float64
	Description string
}

func NewHouseBlend() Beverage {
	return &HouseBlend{
		Price:       HouseBlendPrice,
		Description: "House Blend",
	}
}

func (h *HouseBlend) GetPrice() float64 {
	return h.Price
}

func (h *HouseBlend) GetDescription() string {
	return h.Description
}
