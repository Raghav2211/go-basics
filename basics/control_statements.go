package basics

import (
	"fmt"
	"math"
	"math/rand"
)

func ControlStatements() {

	if true {
		fmt.Println("Always execute")
	}
	if false {
		fmt.Println("Never execute")
	}

	studentGrades := map[string]int{
		"Student1": 12,
		"Student2": 20,
		"Student3": 50,
	}

	if student1, ok := studentGrades["Student1"]; ok {
		fmt.Println("Student1 grade : ", student1)
	}
	// if-else
	var num = 0.1
	if isNumberSameAfterOps(num) {
		fmt.Println("Numbers are same")
	} else {
		fmt.Println("Numbers are different")
	}
	num = 0.123
	if isNumberSameAfterOps(num) {
		fmt.Println("Numbers are same")
	} else {
		fmt.Println("Numbers are different")
	}

	// switch statements
	var generateNumber int = rand.Intn(10-1) + 1
	switch generateNumber {
	case 1:
		fmt.Println("1st Case matched for number", generateNumber)
		//break // break statement are implicits :)
	case 2:
		fmt.Println("2nd Case matched for number", generateNumber)
	default:
		fmt.Println("No Case matched for number", generateNumber)

	}
	generateNumber = rand.Intn(10-1) + 1
	// merge multiple cases
	switch generateNumber {
	case 1, 3, 5, 7, 9:
		fmt.Println("1st Case matched for number", generateNumber)
	case 2, 4, 6, 8:
		fmt.Println("2nd Case matched for number", generateNumber)
	default:
		fmt.Println("No Case matched for number", generateNumber)

	}
	// intialize and compare in switch statements
	switch i := 2 + 3; i {
	case 5:
		fmt.Println("Matched")
	default:
		fmt.Println("No matched")
	}

	// without tag switch statement

	generateNumber = rand.Intn(30-rand.Intn(10-1)+1) + 1
	switch {
	case generateNumber >= 1 && generateNumber <= 10:
		fmt.Println("Generated number between 1-10")
	case generateNumber >= 11 && generateNumber <= 20:
		fmt.Println("Generated number between 11-20")
	default:
		fmt.Println("Generated number above 20")
	}
	// fallthrough in switch statments
	i := 10
	switch {
	case i >= 10:
		fmt.Println("i is greater/equal to 10") // this will print
	case i < 20:
		fmt.Println("i is less than to 20") // this will not print because of implicit break

	}

	switch {
	case i >= 10:
		fmt.Println("i is greater/equal to 10") // this will print
		fallthrough
	case i < 20:
		fmt.Println("i is less than to 20") // this also print becuase of fallthrough

	}

}

func isNumberSameAfterOps(num float64) bool {
	return num == math.Pow(math.Sqrt(num), 2)
}
