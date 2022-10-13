package basics

import (
	"fmt"
)

func Loop() {
	fmt.Println("***************************Loop Section Start***************************")

	// simple for loop
	for i := 0; i < 10; i++ {
		fmt.Println("[Simple loop] i is ", i)
	}
	fmt.Println()
	// mltiple variable
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println("[Multiple variable]i is ", i, " and j is ", j)
	}
	fmt.Println()
	// initialize i in main scope
	var i int = 0
	for ; i < 10; i++ {
		fmt.Println("[Initialize value in main scope] i is ", i)
	}
	fmt.Println()
	// only comparison operator (can compare with while loop)
	i = 0
	for i < 10 {
		fmt.Println("[Only comparison operator] i is ", i)
		i++
	}
	fmt.Println()

	// for loop without expression
	i = 0
	for {
		fmt.Println("[Without expression] i is ", i)
		i++
		if i > 9 {
			break
		}
	}

	fmt.Println()
	// inner loop
	i = 1
	for ; i <= 3; i++ {
		fmt.Println("Table of ", i, " start")
		for j := 1; j <= 10; j++ {
			fmt.Println("", i*j)
		}
		fmt.Println("Table of ", i, " end")
	}

	fmt.Println()

	// break outer loop
	i = 1
Loop:
	for ; i < 10; i++ {
		fmt.Println("Table of ", i, " start")
		for j := 1; j < 10; j++ {

			if i == 5 {
				break Loop
			}
			fmt.Println("", i*j)
		}
		fmt.Println("Table of ", i, " end")
	}
	fmt.Println("Table of ", i, " break")

	fmt.Println()
	// for each loop
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println("index is ", i, " value is ", v)
	}
	fmt.Println()
	studentGrades := map[string]int{
		"Student1": 12,
		"Student2": 20,
		"Student3": 50,
	}
	// both key & value
	for k, v := range studentGrades {
		fmt.Println("[print both key & value]key is ", k, " value is ", v)
	}
	// only key
	for k := range studentGrades {
		fmt.Println("[print only key] key is ", k)
	}
	// only value
	for _, v := range studentGrades {
		fmt.Println("[print only value] value is ", v)
	}
	fmt.Println("***************************Loop Section End***************************")
}
