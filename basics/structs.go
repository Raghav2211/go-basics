package basics

import (
	author "basics/basics/model"
	"fmt"
)

func Struct() {
	fmt.Println("***************************Structs Section Start***************************")
	println()

	authorObj := author.Author{
		Id:        "1",
		FirstName: "Raghav",
		LastName:  "Joshi",
	}
	fmt.Println("Author info: ", authorObj)
	println()

	// there are no build in construtor , so created a new function which works as constructor
	authorObj = author.NewAuthor("Raghav", "Joshi")
	fmt.Println("Author created using constructor : ", authorObj)
	println()

	// check equality of two struct
	copyAuthorObj := authorObj
	fmt.Println("Author created using copy : ", copyAuthorObj)
	fmt.Println("Is actual & copy author equal ? ", (authorObj == copyAuthorObj))

	copyAuthorObjUsingPointer := &copyAuthorObj
	copyAuthorObj.FirstName = "Change Firstname"

	fmt.Println("[After change in copy object]Is actual & copy author equal ? ", (authorObj == copyAuthorObj))
	fmt.Println("Is copy & pointer copy author equal ? ", (*copyAuthorObjUsingPointer == copyAuthorObj))
	println()

	anonymousAuthor := struct {
		id        string
		firstName string
		lastName  string
	}{"1", "Raghav", "Joshi"}
	fmt.Println("Anonymous Author info", anonymousAuthor)
	println()

	// struct assignment
	anonymousAuthor2 := anonymousAuthor

	// update struct attribute will not effect all assignment
	anonymousAuthor2.firstName = "Raghav2"

	fmt.Println("[After update]Anonymous Author info", anonymousAuthor)
	fmt.Println("[After update]Anonymous2 Author info", anonymousAuthor2)
	println()

	// struct assignment using pointer
	refAnonymousAuthor := &anonymousAuthor

	// update struct attribute will  effect all assignment
	refAnonymousAuthor.firstName = "Raghav2"
	fmt.Println("[After update using pointer assignment]Anonymous Author info", anonymousAuthor)
	fmt.Println("[After update using pointer assignment]Reference Anonymous Author info", refAnonymousAuthor)
	println()

	fmt.Println("***************************Structs Section End***************************")

}
