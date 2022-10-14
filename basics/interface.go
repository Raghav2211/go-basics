package basics

import (
	coffeeApp "basics/basics/conffeeapp"
	coffeemachine "basics/basics/conffeeapp/machine"
	"fmt"
)

func Interface() {
	fmt.Println("***************************Interface Section Start***************************")
	coffeeApp := coffeeApp.CoffeeApp{}

	// create basic coffeeMachine
	basicCoffeeBean := coffeemachine.CoffeeBean{Quantity: 100, CoffeeType: coffeemachine.BASIC}
	basicCoffeeMachine := coffeemachine.BasicCoffeeMachine{Bean: basicCoffeeBean}
	// create premium coffee machine for basic type
	premiumBasicCoffeeBean := coffeemachine.CoffeeBean{Quantity: 200, CoffeeType: coffeemachine.BASIC}
	premiumBasicCoffeeMachine := coffeemachine.PremiumCoffeeMachine{Bean: premiumBasicCoffeeBean}
	// create premium coffee machine for espresso type
	premiumEspressoCoffeeBean := coffeemachine.CoffeeBean{Quantity: 1000, CoffeeType: coffeemachine.ESPRESSO}
	premiumEspressoCoffeeMachine := coffeemachine.PremiumCoffeeMachine{Bean: premiumEspressoCoffeeBean}

	coffeeMachines := []coffeemachine.CoffeeMachine{basicCoffeeMachine, premiumBasicCoffeeMachine, premiumEspressoCoffeeMachine}
	for _, coffeeMachine := range coffeeMachines {
		coffeeApp.PrepareCoffee(coffeeMachine)
		println()
	}
	fmt.Println("***************************Interface Section End***************************")
}
