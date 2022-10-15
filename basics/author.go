package basics

import (
	"encoding/json"
	"fmt"
	"log"
)

type Author struct {
	Id        string `json:"id"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
}

func (author Author) String() string {
	authorJSON, err := json.Marshal(author)
	if err != nil {
		log.Fatal("Error occurred during marshalling")
	}
	return string(authorJSON)
}

func (author Author) PrintAuthorName() {
	fmt.Println("Author Name :", author.FirstName, author.LastName)
}
