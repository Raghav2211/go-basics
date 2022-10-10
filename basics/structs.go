package basics

import "fmt"

type Author struct {
	id        string
	firstName string
	lastName  string
}

func Struct() {
	author := Author{
		id:        "1",
		firstName: "Raghav",
		lastName:  "Joshi",
	}

	fmt.Println("Author info: ", author)
	fmt.Println("Author fullName : ", author.firstName+" "+author.lastName)

	anonymousAuthor := struct {
		id        string
		firstName string
		lastName  string
	}{"1", "Raghav", "Joshi"}
	fmt.Println("Anonymous Author info", anonymousAuthor)
	fmt.Println("Anonymous Author info", anonymousAuthor.firstName+" "+anonymousAuthor.lastName)

	anonymousAuthor2 := anonymousAuthor
	anonymousAuthor2.firstName = "new"
	anonymousAuthor2.lastName = "user"

	fmt.Println("Anonymous Author info", anonymousAuthor)
	fmt.Println("Anonymous2 Author info", anonymousAuthor2)

	refAnonymousAuthor := &anonymousAuthor

	refAnonymousAuthor.firstName = "Ref"
	refAnonymousAuthor.lastName = "Change"
	fmt.Println("Anonymous Author info", anonymousAuthor)
	fmt.Println("Reference Anonymous Author info", refAnonymousAuthor)

}
