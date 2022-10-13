package basics

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

func Function() {
	println("***************************Function Section Start***************************")
	// calling a function
	helloWorld()
	println()
	args(1, 2, 3, 4)
	println()
	// calling a function with parameters
	addNumbers(2, 3)
	println()
	// calling function with same parameter does not need to specify each variable type
	addNumbersWithoutDefineTypeofA(2, 3)
	println()
	//passing parameters by value
	name := "Raghav"
	passingParamterByValues(name)
	fmt.Println("Value in caller after function call : ", name)
	println()

	// passing value using pointers
	name = "Raghav"
	passingParamterByPointer(&name)
	fmt.Println("Value in caller after function call : ", name)
	println()

	// update map using function
	studentGrades := map[string]int{
		"Student1": 12,
		"Student2": 14,
		"Student3": 14,
	}
	fmt.Println("before passing to downstream function(updateMap) studentGrades :", studentGrades)
	updateMap(studentGrades)
	fmt.Println("after update studentGrades in downstream function(updateMap) :", studentGrades)
	println()

	// update slice using function
	grades := maps.Values(studentGrades)
	fmt.Println("before passing to downstream function(updateSlice) grades :", grades)
	updateSlice(&grades)
	fmt.Println("after update studentGrades in downstream function(updateSlice) :", grades)
	println()
	// named return function
	nameWithUpperCase := namedResultFunction(name)
	fmt.Println("return uppercase string from named return function :", nameWithUpperCase)
	println()

	// return multiple values from function
	firstname, middlename, lastname := divideFullNameIntoSubname("Raghav Joshi")
	fmt.Println("Return all form of name  with multiple return function [", firstname, middlename, lastname, "]")
	firstname, middlename, lastname = divideFullNameIntoSubname("Raghav")
	fmt.Println("Return all form of name  with multiple return function [", firstname, middlename, lastname, "]")
	fmt.Println()
	println("***************************Function Section End***************************")
}

func helloWorld() { fmt.Println("Hello world") }

func args(values ...int) {
	for _, value := range values {
		fmt.Print(value, ",")
	}
	println()
}
func addNumbers(a int, b int) {
	fmt.Println("a + b : ", (a + b))
}

func addNumbersWithoutDefineTypeofA(a, b int) {
	fmt.Println("a + b : ", (a + b))
}

func passingParamterByValues(name string) {
	fmt.Println("Value recieved : ", name)
	name = "Change"
	fmt.Println("After Value change in callee : ", name)
}

func passingParamterByPointer(name *string) {
	fmt.Println("Value recieved : ", *name)
	*name = "Change"
	fmt.Println("After Value change in callee : ", *name)
}

func updateMap(studentGrades map[string]int) {
	studentGrades["Student2"] = 13
}
func updateSlice(grades *[]int) {
	*grades = append(*grades, 15)
}

func namedResultFunction(name string) (result string) {
	result = strings.ToUpper(name)
	return
}

func divideFullNameIntoSubname(fullname string) (string, string, string) {
	nameArray := strings.Split(fullname, " ")
	if len(nameArray) == 1 {
		return nameArray[0], "", ""
	}
	if len(nameArray) == 2 {
		return nameArray[0], "", nameArray[1]
	}
	return nameArray[0], nameArray[1], nameArray[2]
}
