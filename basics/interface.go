package basics

import (
	coffeeApp "basics/basics/conffeeapp"
	coffeemachine "basics/basics/conffeeapp/machine"
	"fmt"
)

type struct1 struct {
	struct1Name string
}
type struct2 struct {
	struct2Name string
}

func Interface() {
	fmt.Println("***************************Interface Section Start***************************")
	// empty interface
	var emptyInterface interface{}

	emptyInterface = 12
	describe(emptyInterface)
	println()

	emptyInterface = "Raghav"
	describe(emptyInterface)
	println()

	// type assertion -- type assertion provides access to an interface value's underlying concrete value.
	var structtype interface{} = struct1{"struct1"}
	describeAndPrintName(structtype)
	println()

	structtype = struct2{"struct2"}
	describeAndPrintName(structtype)
	println()

	fmt.Println("------------------------Coffee Example------------------------")
	// coffee example
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

func describe(value interface{}) {
	fmt.Printf("%v % T", value, value)
}

func describeAndPrintName(structType interface{}) {
	switch x := structType.(type) {
	case struct1:
		fmt.Printf("(%v %T)", x.struct1Name, x)
	case struct2:
		fmt.Printf("(%v %T)", x.struct2Name, x)

	}

}
