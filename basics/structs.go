package basics

import "fmt"

func Struct() {
	fmt.Println("***************************Structs Section Start***************************")

	author := Author{
		Id:        "1",
		FirstName: "Raghav",
		LastName:  "Joshi",
	}

	fmt.Println("Author info: ", author)

	anonymousAuthor := struct {
		id        string
		firstName string
		lastName  string
	}{"1", "Raghav", "Joshi"}
	fmt.Println("Anonymous Author info", anonymousAuthor)

	// struct assignment
	anonymousAuthor2 := anonymousAuthor

	// update struct attribute will not effect all assignment
	anonymousAuthor2.firstName = "Raghav2"

	fmt.Println("[After update]Anonymous Author info", anonymousAuthor)
	fmt.Println("[After update]Anonymous2 Author info", anonymousAuthor2)

	// struct assignment using pointer
	refAnonymousAuthor := &anonymousAuthor

	// update struct attribute will  effect all assignment
	refAnonymousAuthor.firstName = "Raghav2"
	fmt.Println("[After update using pointer assignment]Anonymous Author info", anonymousAuthor)
	fmt.Println("[After update using pointer assignment]Reference Anonymous Author info", refAnonymousAuthor)

	fmt.Println("***************************Structs Section End***************************")

}
