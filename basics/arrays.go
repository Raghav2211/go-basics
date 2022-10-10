package basics

import "fmt"

// fixed in size cannot grow after initilaization
func Arrays() {
	// declaration styles
	grades := [3]int{1, 2, 3}
	fmt.Printf("Grades %v\n", grades)

	// without decalre size
	gradesWithoutSize := [...]int{1, 2, 3, 4}
	fmt.Printf("Grades Without size declaration %v\n", gradesWithoutSize)

	var students [3]string
	//[3]string{"Grabage"} // initialization
	fmt.Printf("Students Empty array %v\n", students)
	students[0] = "Studetn1"
	fmt.Printf("Students %v\n", students)
	students[1] = "Student2"
	students[2] = "Student3"
	fmt.Printf("Students %v\n", students)
	// students[3] = "Student4" // out of bound

	fmt.Println("Length of students array", len(students))

	// multi dimensional
	var matrix [3][3]int = [3][3]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}
	fmt.Printf("Matrix %v\n", matrix)

	// copy one array to another
	a := [...]int{1, 2, 3}
	b := a // copy data (don't use , use pointers)
	b[1] = 4
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 4
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", d[0])

}
