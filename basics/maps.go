package basics

import (
	author "basics/basics/model"
	"fmt"
)

func Maps() {
	fmt.Println("***************************Maps Section Start***************************")

	// declaration
	var emptyMap = map[string]int{}

	fmt.Println(emptyMap)
	println()

	// using make built-in function
	var makeMap = make(map[string]int)
	fmt.Println(makeMap)
	println()

	studentGrades := map[string]int{
		"Student1": 12,
		"Student2": 20,
		"Student3": 50,
	}
	fmt.Println(studentGrades)
	println()

	studentGrades["Student4"] = 15 // add

	fmt.Println(studentGrades)
	fmt.Println("Student1 grade =", studentGrades["Student1"]) // retrieve
	fmt.Println("Student5 grade =", studentGrades["Student5"]) // exists or not ?
	println()

	delete(studentGrades, "Student1") // delete
	fmt.Println("Grab all student grades", studentGrades)
	println()

	// sure key is present or not
	student1, ok1 := studentGrades["Student1"]
	student2, ok2 := studentGrades["Student2"]
	fmt.Println(student1, ok1)
	fmt.Println(student2, ok2)
	println()

	// or just need to check key is present or not

	_, ok3 := studentGrades["Student1"]
	_, ok4 := studentGrades["Student2"]
	fmt.Println("is Student1 present?", ok3)
	fmt.Println("is Student2 present?", ok4)
	fmt.Println("Length of studentGrade db ", len(studentGrades))
	println()

	// multiple map assignment use same underline data
	map1 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	map2 := map1
	fmt.Println("Map1 : ", map1)
	fmt.Println("Map2 :", map2)
	map2["key4"] = 4 // append
	fmt.Println("After append in map2 , Map1: ", map1)
	fmt.Println("After append in map2 , Map2: ", map2)
	delete(map1, "key1")
	fmt.Println("After delete[key1] in map1 , Map1: ", map1)
	fmt.Println("After delete[key1] in map1 , Map2: ", map2)
	println()

	// composite key for map
	authorObj := author.NewAuthor("1FName", "1LName")

	authorObj2 := authorObj
	authorObj3 := author.NewAuthor("3FName", "2LName")

	authorGrades := map[author.Author]int{
		authorObj:  12,
		authorObj2: 13,
		authorObj3: 14,
	}
	fmt.Println("build map with composite key", authorGrades) // should contain 2 objects as authorObj & authorObj2 are equal

	fmt.Println("***************************Primitives Section End***************************")

}
