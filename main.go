package main

import (
	"decorator-pattern/beverage"
	"decorator-pattern/condiment"
	"fmt"
)

func main() {
	beverage1 := beverage.NewDarkRoast()
	beverage1 = condiment.NewMilk(beverage1)
	beverage1 = condiment.NewSoy(beverage1)

	fmt.Println(beverage1.GetPrice())
	fmt.Println(beverage1.GetDescription())
}
