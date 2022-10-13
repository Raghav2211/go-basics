package basics

import "fmt"

func Pointers() {

	println("***************************Pointers Section Start***************************")
	a := 1
	b := a // copy a variable data into b (a & b are not sharing the same memory)
	println("value of a is ", a, ", value of b is ", b)
	println("address of a is ", &a, ", address of b is ", &b) // this should different
	a = 2
	println(a, b)
	println()
	// if we want to share the memory location then
	var c int = 2
	var d *int = &c
	println("value of c is ", c, ", value of d is ", *d)     // dereferencing of pointer d
	println("address of c is ", &c, ", address of d is ", d) // this should same
	println()
	// if change c value it should affect on d also
	c = 3
	println("value of c is ", c, ", value of d is ", *d)
	println("address of c is ", &c, ", address of d is ", d)

	println()
	// if change d value it should affect on c also
	*d = 4
	println("value of c is ", c, ", value of d is ", *d)
	println("address of c is ", &c, ", address of d is ", d)

	println()
	// array with pointers
	var customArray = [3]int{1, 3, 3}
	copyArray := customArray
	fmt.Println("custom array values: ", customArray, " copy array values: ", copyArray)
	customArray[1] = 2                                                                                 // will not affect the copyArray
	fmt.Println("[After change]custom array values: ", customArray, " copy array values: ", copyArray) // this should be different becuase array does copy it's content

	referenceArray := &customArray
	copyArray[1] = 3
	fmt.Println("[Using pointer]custom array values: ", customArray, " reference array values: ", *referenceArray) //  both should be same

	println()

	// structs with pointers
	customStruct := customStruct{12}
	copyStruct := customStruct
	fmt.Println("custom struct value : ", customStruct, " copyStruct struct value : ", copyStruct)
	customStruct.valueInInt = 13 // does not effect of copyStruct
	fmt.Println("[After change]custom struct value : ", customStruct, " copyStruct struct value : ", copyStruct)
	referenceStruct := &customStruct
	fmt.Println("custom struct value : ", customStruct, " referenceStruct struct value : ", *referenceStruct)
	customStruct.valueInInt = 14 // will effect referenceStruct
	fmt.Println("[After change]custom struct value : ", customStruct, " referenceStruct struct value : ", *referenceStruct)

	// slices & maps does not copy the underline data it copy the refernece

	println("***************************Pointers Section End***************************")
}

type customStruct struct {
	valueInInt int
}
