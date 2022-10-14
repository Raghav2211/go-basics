package machine

import "fmt"

type CoffeeType int

const (
	BASIC CoffeeType = iota
	ESPRESSO
)

type CoffeeMachine interface {
	BrewCoffee() string
}

type EspressoCoffeeMachine interface {
	CoffeeMachine
	BrewEspresso() string
}

type grinder struct {
}

type CoffeeBean struct {
	Quantity   int
	CoffeeType CoffeeType
}
type BasicCoffeeMachine struct {
	Bean CoffeeBean
}
type PremiumCoffeeMachine struct {
	Bean CoffeeBean
}

func (basicCoffeeMachine BasicCoffeeMachine) BrewCoffee() string {
	grinder := grinder{}

	grinder.grind(basicCoffeeMachine.Bean)
	return fmt.Sprintf("filter coffee ready !!")
}

func (premiumCoffeeMachine PremiumCoffeeMachine) BrewEspresso() string {
	return fmt.Sprintf("espresso coffee ready !!")
}

func (premiumCoffeeMachine PremiumCoffeeMachine) BrewCoffee() string {
	grinder := grinder{}

	grinder.grind(premiumCoffeeMachine.Bean)
	if premiumCoffeeMachine.Bean.CoffeeType == BASIC {
		return fmt.Sprintf("filter coffee ready !!")
	}
	return premiumCoffeeMachine.BrewEspresso()
}

func (grinder grinder) grind(coffeeBean CoffeeBean) {
	fmt.Println("Grind coffee with quantity : ", coffeeBean.Quantity)
}
