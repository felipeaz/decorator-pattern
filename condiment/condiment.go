package condiment

const (
	MilkPrice  = 1.00
	MochaPrice = 1.29
	SoyPrice   = 1.15
	WhipPrice  = 1.45
)

type Condiment interface {
	GetDescription()
}
