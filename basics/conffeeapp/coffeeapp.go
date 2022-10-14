package conffeeapp

import (
	machine "basics/basics/conffeeapp/machine"
	"fmt"
)

type CoffeeApp struct {
}

func (coffeeApp CoffeeApp) PrepareCoffee(coffeeMachine machine.CoffeeMachine) {
	fmt.Println(coffeeMachine.BrewCoffee())
}
