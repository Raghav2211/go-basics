package basics

import "fmt"

func Maps() {
	// declaration
	var emptyMap = map[string]int{}
	emptyMap2 := map[string]int{}

	fmt.Println(emptyMap)
	fmt.Println(emptyMap2)

	var makeMap = make(map[string]int)
	fmt.Println(makeMap)

	studentGrades := map[string]int{
		"Student1": 12,
		"Student2": 20,
		"Student3": 50,
	}
	fmt.Println(studentGrades)

	studentGrades["Student4"] = 15 // add

	fmt.Println(studentGrades)
	fmt.Println("Student1 grade =", studentGrades["Student1"]) // retrieve
	fmt.Println("Student5 grade =", studentGrades["Student5"]) // exists or not ?

	delete(studentGrades, "Student1") // delete
	fmt.Println("Grab all student grades", studentGrades)

	// sure key is present or not
	student1, ok1 := studentGrades["Student1"]
	student2, ok2 := studentGrades["Student2"]
	fmt.Println(student1, ok1)
	fmt.Println(student2, ok2)

	// or just need to check key is present or not

	_, ok3 := studentGrades["Student1"]
	_, ok4 := studentGrades["Student2"]
	fmt.Println("is Student1 present?", ok3)
	fmt.Println("is Student2 present?", ok4)
	fmt.Println("Length of studentGrade db ", len(studentGrades))

}
