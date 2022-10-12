package basics

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Defer() {

	fmt.Println("***************************Defer Section Start***************************")
	// defer calls goes into stack and run in LIFO order
	defer fmt.Println("***************************Defer Section End***************************")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // despite panic this would run
	}

	resp, err := http.Get("https://run.mocky.io/v3/6b20f760-08c2-4694-9109-ac369d2d2a8e")
	if err != nil {
		panic("Error occurred")
	}
	body := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&body)
	fmt.Println("Got respose : ", body, " with responseCode : ", resp.StatusCode)
	defer resp.Body.Close()
}
