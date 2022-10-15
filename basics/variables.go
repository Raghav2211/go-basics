package basics

import (
	"fmt"
	"strconv"
)

// package level variables

var k int32 = 43

// local variable hide this variable
var shadow int32 = 1

var (
	name string = "raghav"
	age  int32  = 30
)

var I int32 = 4 // GLOBAL variable

func Variables() {
	fmt.Println("***************************Variables Section Start***************************")
	// local variables
	var i int32 = 35
	j := 36 // := is walrus operator
	fmt.Printf("[int]Local variable i %v , %T\n", i, i)
	fmt.Printf("[int]Local variable j %v , %T\n", j, j)
	fmt.Printf("[int]Package level variable K %v , %T\n", k, k)

	var shadow int32 = 2
	fmt.Printf("[int]Local variable shadow %v , %T\n", shadow, shadow)

	fmt.Printf("[string]Package variable name %v , %T\n", name, name)

	var floatVar float32 = 32.5
	var intVar int32 = int32(floatVar)
	fmt.Printf("[float]Local variable floatVar %v , %T\n", floatVar, floatVar)
	fmt.Printf("[int]Local variable intVar %v , %T\n", intVar, intVar)

	var stringVar string = string(intVar)
	fmt.Printf("[String]Local variable stringVar %v , %T\n", stringVar, stringVar)

	var stringToIntConvert string = strconv.Itoa(int(intVar))
	fmt.Printf("Local variable %v , %T\n", stringToIntConvert, stringToIntConvert)

	fmt.Println("***************************Variables Section End***************************")

}
