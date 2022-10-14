package conffeeapp

import (
	coffeemachine "basics/basics/conffeeapp/machine"
	"fmt"
)

type CoffeeApp struct {
}

func (coffeeApp CoffeeApp) PrepareCoffee(coffeeMachine coffeemachine.CoffeeMachine) {
	describeType(coffeeMachine)
	fmt.Println(coffeeMachine.BrewCoffee())
}

func describeType(coffeeMachine coffeemachine.CoffeeMachine) {
	fmt.Printf("Start machine %T\n ", coffeeMachine) // interface value holds a value of a specific underlying concrete type.
}
